package main

import (
	"fmt"
	"strconv"
)

type Computer interface {
	addRAM(int)
	toString() string
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
			inside := PC{name: "desktop" + strconv.Itoa(i+1), ram: (i + 1) * 1024, cpu: "intel " + strconv.Itoa(i+2) + "86DX",
				gpu: "AMD" + strconv.Itoa(i+1000) + "HD", cost: (i + 1) + (i+1)*300, storage: (i+1)*1000 + i*2*8*16}
			container.computers[i] = Laptop{brand: "Brand" + strconv.Itoa(i), isBacklit: true, screen: 15.6, weight: 15.0*float32(i) + 1000.0, inner: inside}

		} else {
			container.computers[i] = PC{name: "desktop" + strconv.Itoa(i+1), ram: (i + 1) * 1024, cpu: "intel " + strconv.Itoa(i+2) + "86DX",
				gpu: "AMD" + strconv.Itoa(i+1000) + "HD", cost: (i + 1) + (i+1)*300, storage: (i+1)*1000 + i*2*8*16}
		}
	}
	for i, v := range container.computers {
		fmt.Println(strconv.Itoa(i) + " | " + v.toString())
	}
}

func (object PC) addRAM(additionRAM int) {
	object.ram += additionRAM
}

func (object PC) toString() string {
	return fmt.Sprintf("PC!\t\t name: %v, cpu: %v, gpu: %v, storage: %v, ram: %v, cost: %v", object.name, object.cpu, object.gpu, object.storage, object.ram, object.cost)
}

func (object Laptop) addRAM(additionRAM int) {
	object.inner.ram += additionRAM
}

func (object Laptop) toString() string {
	return fmt.Sprintf("Laptop!\t brand: %v, weight: %v, screen: %v, isBacklit: %v, name: %v, cost: %v",
		object.brand, object.weight, object.screen, object.isBacklit, object.inner.name, object.inner.cost)
}
