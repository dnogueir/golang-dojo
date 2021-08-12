package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		groupCounter("pipoca", 10)
		
		wg.Done()
	}()
	go func() {
		groupCounter("netflix", 2)
		wg.Done()
	}()

	wg.Wait()

}

func groupCounter(something string, value int) {
	for i := 0; i < value; i++ {
		fmt.Println(something)
		time.Sleep(time.Millisecond * 500)
	}
}
