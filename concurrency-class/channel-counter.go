package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)

	go channelCounter("pipoca", c)
	go channelCounter("netflix", c)

	for msg := range c {
		fmt.Println(msg)
	}

	//close(c)

}

func channelCounter(something string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- something
		time.Sleep(time.Millisecond * 500)
	}

}
