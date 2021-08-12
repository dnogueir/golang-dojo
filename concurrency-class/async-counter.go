package main

import (
	"fmt"
	"time"
)

//go routine
func main() {

	go asyncCounter("pipoca", 15)
	go asyncCounter("netflix", 2)

}

func asyncCounter(something string, valor int) {
	for i := 0; i < valor; i++ {
		fmt.Println(something)
		time.Sleep(time.Millisecond * 500)
	}
}
