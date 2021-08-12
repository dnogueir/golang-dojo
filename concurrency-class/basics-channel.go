package main

import "fmt"

func main() {

	//firstExample()
	//secondExample()
	//thirdExample()
}

func firstExample() {
	c := make(chan string)
	c <- "first example"

	fmt.Println(<-c)
}

func secondExample() {

	c := make(chan string, 1)
	c <- "second example"

	fmt.Println(<-c)
}

func thirdExample() {
	c := make(chan string)

	go func() {
		c <- "third example"
	}()

	fmt.Println(<-c)

}
