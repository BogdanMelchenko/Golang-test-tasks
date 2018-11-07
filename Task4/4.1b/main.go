package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	once := sync.Once{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			once.Do(func() {
				defer fmt.Println("world")
				fmt.Println("hello")
			})
		}()
	}
	wg.Wait()
	fmt.Scanln()
}
