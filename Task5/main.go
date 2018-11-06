package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"net/http"
	"sort"
	"time"
)

func ping(index int, adress string, client http.Client, resultChannel chan float64) {
	resp, err := client.Get(adress)
	if err != nil && net.Error.Timeout(err.(net.Error)) {
		resultChannel <- -1.0
		return
	}
	start := time.Now()
	var duration float64
	defer resp.Body.Close()
	for {
		bs := make([]byte, 1014)
		_, err := resp.Body.Read(bs)
		if err != nil {
			finish := time.Now()
			elapsed := finish.Sub(start)
			duration = float64(elapsed) / float64(time.Millisecond)
			resultChannel <- duration
			break
		}
	}
}

func main() {

	quantity := flag.Int("quantity", 10, "quantity of requests")
	var address string
	flag.StringVar(&address, "address", "https://google.com", "where to send requests")
	timeout := flag.Int("timeout", 6, "timeout of request")
	flag.Parse()

	resultChannel := make(chan float64)
	timeoutDuration := time.Duration(*timeout)
	client := http.Client{
		Timeout: timeoutDuration * time.Second,
	}

	startTime := time.Now()
	for i := 0; i < *quantity; {
		go ping(i, address, client, resultChannel)
		i++
	}
	i := 0
	timeouts := 0
	results := make([]float64, *quantity)
	for i < *quantity {
		select {
		case result := <-resultChannel:
			if result < float64(0.0) {
				timeouts++
			}
			results[i] = result
			i++
		default:
			continue
		}
	}
	totalTime := time.Now().Sub(startTime)

	filteredResults := results[:0]
	for _, res := range results {
		if res > 0.0 {
			filteredResults = append(filteredResults, res)
		}
	}
	sort.Float64s(filteredResults)

	min := filteredResults[0]
	max := filteredResults[len(filteredResults)-1]

	printResults(address, min, max, timeouts, *quantity, totalTime)
	fmt.Scanln()
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFloat64Milliseconds(goal time.Duration) float64 {
	return float64(goal) / float64(time.Millisecond)
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func printResults(address string, min, max float64, timeouts, quantity int, totalTime time.Duration) {
	fmt.Println("total time for", address, "is:", toFixed(toFloat64Milliseconds(totalTime), 2), "ms")
	fmt.Println("average is: ", toFixed(toFloat64Milliseconds(totalTime)/float64(quantity), 2), "ms")
	fmt.Println("request timeouts / total requests:", timeouts, "/", quantity)
	fmt.Println("min is:", toFixed(min, 2), "ms")
	fmt.Println("max is:", toFixed(max, 2), "ms")
}
