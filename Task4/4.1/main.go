package main

import (
	"fmt"
	"sync"
)

func helloWorld() func() {
	return func() {
		defer fmt.Println(" World")
		fmt.Println(" Hello")
	}
}

func main() {

	f := helloWorld()
	once := sync.Once{}
	onceBody := func() {
		f()
	}

	done := make(chan bool)
	for i := 0; i < 10; {
		go func() {
			once.Do(onceBody)
			done <- true
		}()

		i++
	}

	<-done
	fmt.Scanln()
}
