package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	single_threaded()

	multi_threaded()
}

func multi_threaded() {
	startTime := time.Now()
	var list []int
	var wg sync.WaitGroup

	slice_count := 8
	list = generate_slice(slice_count * 100)

	// c := make(chan int, 8)
	c := make(chan int)

	wg.Add(slice_count)
	for i := 0; i < slice_count; i++ {
		slice_size := i * 100

		go func() {
			c <- sum_values(sub_slice(list, slice_size, slice_size+100))
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	sum := 0
	for partial_sum := range c {
		sum += partial_sum
	}

	fmt.Println(sum)
	fmt.Println(time.Since(startTime))
}

func single_threaded() {
	startTime := time.Now()
	var list []int

	slice_count := 8
	list = generate_slice(slice_count * 100)

	fmt.Println(sum_values(list))
	fmt.Println(time.Since(startTime))
}

func sum_values(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func generate_slice(value int) []int {
	var list []int
	for i := 0; i < value; i++ {
		list = append(list, i)
	}
	return list
}

func sub_slice(slice []int, start int, end int) []int {
	var sub_slice []int
	for i := start; i < end; i++ {
		sub_slice = append(sub_slice, slice[i])
	}
	return sub_slice
}
