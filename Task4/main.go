package main

import (
	"fmt"
)

func helloWorld() func() {
	return func() {
		defer fmt.Println(" World")
		fmt.Println(" Hello")
	}
}

func main() {
	f := helloWorld()
	f()

	fmt.Scanln()
}
