package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		groupCounter("pipoca")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		groupCounter("netflix")
		wg.Done()
	}()

	wg.Wait()

}

func groupCounter(something string) {
	for i := 0; i < 5; i++ {
		fmt.Println(something)
		time.Sleep(time.Millisecond * 500)
	}
}
