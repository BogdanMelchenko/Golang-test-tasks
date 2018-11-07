package main

import (
	"fmt"
	"sync"
)

func main() {
	once := sync.Once{}

	done := make(chan bool)
	for i := 0; i < 10; {
		go func() {
			once.Do(
				func() {
					defer fmt.Println(" World")
					fmt.Println(" Hello")
				})
			done <- true
		}()

		i++
	}

	<-done
	fmt.Scanln()
}
