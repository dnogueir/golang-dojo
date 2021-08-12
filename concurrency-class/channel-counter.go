package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan string)

	wg.Add(2)

	go channelCounter("pipoca", c)
	go channelCounter("netflix", c)

	for msg := range c {
		fmt.Println(msg)
	}

	wg.Wait()
	close(c)

}

func channelCounter(something string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- something
		time.Sleep(time.Millisecond * 500)
	}

	wg.Done()

}
