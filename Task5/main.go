package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"sort"
	"time"
)

// Ping returns request duration or -1.0 (float64) if gets Timeout error
func Ping(adress string, client http.Client, resultChannel chan float64) {
	startTime := time.Now()
	response, err := client.Get(adress)
	if err != nil {
		resultChannel <- -1.0
		return
	}
	var durationInFloat float64
	defer response.Body.Close()
	elapsedOnRequest := time.Now().Sub(startTime)
	durationInFloat = float64(elapsedOnRequest) / float64(time.Millisecond)
	resultChannel <- durationInFloat
}

func main() {
	requestQuantity := flag.Int("quantity", 10, "quantity of requests")
	var requestAddress string
	flag.StringVar(&requestAddress, "address", "https://google.com", "where to send requests")
	timeout := flag.Int("timeout", int(1000), "timeout of request in ms")
	flag.Parse()
	fmt.Println("request adress: ", requestAddress, "|\tquantity of requests: ", *requestQuantity, "|\trequest timeout", *timeout, "ms")
	resultChannel := make(chan float64)

	i := time.Duration(*timeout)
	timeoutDuration := time.Duration(i * time.Millisecond)
	client := http.Client{
		Timeout: timeoutDuration,
	}

	startTime := time.Now()
	for i := 0; i < *requestQuantity; {
		go Ping(requestAddress, client, resultChannel)
		i++
	}

	timeoutsQuantity := 0
	results := make([]float64, *requestQuantity)
	for i := 0; i < *requestQuantity; {
		select {
		case result := <-resultChannel:
			if result < float64(0.0) {
				timeoutsQuantity++
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
		if res > -1.0 {
			filteredResults = append(filteredResults, res)
		}
	}

	if len(filteredResults) > 0 {
		sort.Float64s(filteredResults)
		minResponseTime := filteredResults[0]
		maxResponseTime := filteredResults[len(filteredResults)-1]
		averageResponseTime := calculateAverage(filteredResults)
		printResults(requestAddress, minResponseTime, maxResponseTime, averageResponseTime, timeoutsQuantity, *requestQuantity, totalTime)
	} else {
		fmt.Println("no response, try changing request timeout value")
	}
	fmt.Scanln()
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// DurationToFloat64 makes float value from time.duration (returns float64 value of milliseconds)
func DurationToFloat64(goal time.Duration) float64 {
	return float64(goal) / float64(time.Millisecond)
}

func calculateAverage(target []float64) float64 {
	sum := 0.0
	for _, value := range target {
		sum += value
	}
	return sum / float64(len(target))
}

// toFixed truncates float value to given precision
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func printResults(requestAddress string, minResponseTime, maxResponseTime, averageResponseTime float64, timeoutsQuantity, requestQuantity int, totalTime time.Duration) {
	fmt.Println("total time for", requestAddress, "is:", toFixed(DurationToFloat64(totalTime), 2), "ms")
	fmt.Println("average is: ", toFixed(averageResponseTime, 2), "ms")
	fmt.Println("request timeouts / total requests:", timeoutsQuantity, "/", requestQuantity)
	fmt.Println("min is:", toFixed(minResponseTime, 2), "ms")
	fmt.Println("max is:", toFixed(maxResponseTime, 2), "ms")
}
