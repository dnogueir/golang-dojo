package main

import (
	"fmt"
	"sync"
)

func main() {
	var list []int
	var wg sync.WaitGroup

	list = generate_slice(800)

	c := make(chan int)

	wg.Add(8)
	go func() {
		c <- sum_values(sub_slice(list, 0, 100))
		wg.Done()
	}()
	go func() {
		c <- sum_values(sub_slice(list, 100, 200))
		wg.Done()
	}()
	go func() {
		c <- sum_values(sub_slice(list, 200, 300))
		wg.Done()
	}()
	go func() {
		c <- sum_values(sub_slice(list, 300, 400))
		wg.Done()
	}()
	go func() {
		c <- sum_values(sub_slice(list, 400, 500))
		wg.Done()
	}()
	go func() {
		c <- sum_values(sub_slice(list, 500, 600))
		wg.Done()
	}()
	go func() {
		c <- sum_values(sub_slice(list, 600, 700))
		wg.Done()
	}()
	go func() {
		c <- sum_values(sub_slice(list, 700, 800))
		wg.Done()
	}()

	sum := 0
	for partial_sum := range c {
		sum += partial_sum
	}

	fmt.Println(sum)
	wg.Wait()
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
