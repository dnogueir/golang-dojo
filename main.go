package main

import "fmt"

func main() {
	var list []int

	list = append(list, 1, 2, 3)

	fmt.Println(sum_values(list))
}

func sum_values(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
