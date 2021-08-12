package main

import "fmt"

func main() {
	var list []int
	var sub_list []int

	list = generate_slice(800)

	sub_list = sub_slice(list, 0, 100)
	
	
	fmt.Println(sum_values(sub_list))
	fmt.Println(sum_values(list))
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



