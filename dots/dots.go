package main

import (
	"fmt"
	"time"
)

func writeDots() {
	for {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
}

func main() {
	go writeDots()
	time.Sleep(time.Second * 5)
	fmt.Println("\ndone")
}
