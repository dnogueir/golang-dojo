package main

import "fmt"

func main() {

	//firstExample()
	//secondExample()
	thirdExample()
}

func firstExample() {
	c := make(chan string)
	c <- "first example"
	msg := <- c 
	fmt.Println(msg)
}

func secondExample() {

	c := make(chan string, 1)
	c <- "second example"
	msg := <- c 
	fmt.Println(msg)
}

func thirdExample() {
	c := make(chan string)

	go func() {
		c <- "third example"
	}()
	msg := <- c 
	fmt.Println(msg)

}
