package main

import (
	"fmt"
)

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

func main() {
	desktopOne := PC{gpu: "RTX 3070", cpu: "intel 8086", ram: 4096, storage: 1024, cost: 999, name: "desktopOne"}
	fmt.Println("ram on "+desktopOne.name+" before addition:", desktopOne.ram)
	desktopOne.addRAM(1024 * 2)

	fmt.Println("ram on "+desktopOne.name+" after addition:", desktopOne.ram)
	fmt.Println("============")
	fmt.Println("is "+desktopOne.name+" cost too much?", isCostTooMuch(desktopOne))
	fmt.Println("============")
	laptopOne := Laptop{weight: 1036.6, screen: 15.6, brand: "Lenivo", isBacklit: true}
	laptopOne.inner.ram = 2048
	laptopOne.inner.cost = 3999
	laptopOne.inner.cpu = "intel 286DX"
	laptopOne.inner.gpu = "VooDoo 2"
	laptopOne.inner.name = "laptopOne"

	fmt.Println("ram on "+laptopOne.inner.name+" before addition:", laptopOne.inner.ram)
	laptopOne.inner.addRAM(2048)
	fmt.Println("ram on "+laptopOne.inner.name+" after addition:", laptopOne.inner.ram)
	fmt.Println("============")
	fmt.Println("is "+laptopOne.inner.name+" cost too much?", isCostTooMuch(laptopOne.inner))
}

func isCostTooMuch(object PC) bool {
	if object.cost > 1000 {
		return true
	}
	return false
}

func (object *PC) addRAM(additionRAM int) {
	object.ram += additionRAM
}
