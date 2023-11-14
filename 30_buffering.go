package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "Hello"
	messages <- "World"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
