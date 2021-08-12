package main

import (
	"fmt"
	"time"
)

func main() {

	go asyncCounter("pipoca")
	go asyncCounter("netflix")

}

func asyncCounter(something string) {
	for i := 0; i < 5; i++ {
		fmt.Println(something)
		time.Sleep(time.Millisecond * 500)
	}
}
