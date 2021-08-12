package main

import (
	"fmt"
	"time"
)

func main() {

	syncCounter("pipoca")
	syncCounter("netflix")

}

func syncCounter(something string) {
	for i := 0; i < 5; i++ {
		fmt.Println(something)
		time.Sleep(time.Millisecond * 500)
	}
}
