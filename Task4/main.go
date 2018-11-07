package main

import (
	"fmt"
)

func main() {
	func() {
		defer fmt.Println(" World")
		fmt.Println(" Hello")
	}()

	fmt.Scanln()
}
