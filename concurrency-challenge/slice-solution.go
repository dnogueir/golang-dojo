package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	cSum := make(chan int, 10)

	var totalSum int
	var slice []int

	fillSlice(&slice)

	wg.Add(8)
	start := time.Now()
	go sumSlice(slice[0:100], cSum)
	go sumSlice(slice[100:200], cSum)
	go sumSlice(slice[200:300], cSum)
	go sumSlice(slice[300:400], cSum)
	go sumSlice(slice[400:500], cSum)
	go sumSlice(slice[500:600], cSum)
	go sumSlice(slice[600:700], cSum)
	go sumSlice(slice[700:800], cSum)

	wg.Wait()
	close(cSum)

	for sum := range cSum {
		totalSum += sum
	}

	fmt.Println("paralel sum: ", totalSum)
	fmt.Println("total time:", time.Since(start))

	var totalSync int
	start = time.Now()
	for j := 0; j < 800; j++ {
		totalSync += j
	}

	fmt.Println("Sync Sum:", totalSync)
	fmt.Println("total time:", time.Since(start))

}

func sumSlice(slice []int, cSum chan int) {
	var sum int
	for _, v := range slice {
		sum += v
	}
	cSum <- sum
	wg.Done()
}

func fillSlice(slice *[]int) {
	for i := 0; i < 800; i++ {
		*slice = append(*slice, int(i))
	}
}
