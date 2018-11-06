package main

import (
	"fmt"
	"strconv"
)

type PC struct {
	name    string
	gpu     string
	cpu     string
	ram     int
	storage int
	cost    int
}

func main() {
	var desktops [5]PC
	for i := 0; i < 5; i++ {
		desktops[i].name = "desktop" + strconv.Itoa(i+1)
		desktops[i].ram = (i + 1) * 1024
		desktops[i].cpu = "intel " + strconv.Itoa(i+2) + "86DX"
		desktops[i].gpu = "AMD" + strconv.Itoa(i+1000) + "HD"
		desktops[i].cost = (i + 1) + (i+1)*300
		desktops[i].storage = (i+1)*1000 + i*2*8*16
	}
	for i, v := range desktops {
		fmt.Print(strconv.Itoa(i) + " | ")
		fmt.Println(v)
	}
}
