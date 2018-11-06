package main

import (
	"fmt"
	"strconv"
)

type Computer interface {
	addRAM(int)
	showRAM() string
}

type PC struct {
	name    string
	gpu     string
	cpu     string
	ram     int
	storage int
	cost    int
}

type Laptop struct {
	inner     PC
	weight    float32
	screen    float32
	brand     string
	isBacklit bool
}

type ComputerContainer struct {
	computers [5]Computer
}

func main() {
	var container ComputerContainer
	for i := 0; i < 5; i++ {
		if i%2 == 1 {
			container.computers[i] = new(Laptop)
		} else {
			container.computers[i] = new(PC)
		}
	}
	for i, v := range container.computers {
		v.addRAM(512 * i)
		fmt.Println(strconv.Itoa(i) + " | " + v.showRAM())
	}
}

func (object *PC) addRAM(additionRAM int) {
	object.ram += additionRAM
}

func (object PC) showRAM() string {
	return fmt.Sprintf("PC!\t\t  ram: %v", object.ram)
}

func (object *Laptop) addRAM(additionRAM int) {
	object.inner.ram += additionRAM
}

func (object Laptop) showRAM() string {
	return fmt.Sprintf("Laptop!\t ram: %v", object.inner.ram)
}
