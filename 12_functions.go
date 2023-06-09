package main

import "fmt"

func main() {
	// add
	result := add(7, 10)
	fmt.Println("add : ", result)

	sayHello()
	message := "Go"
	sayMessage(message)

	sayMessageWithPointer(&message)
	fmt.Println(message)
}

func add(a int, b int) int {
	return a + b
}

func sayHello() {
	fmt.Println("sayHello : Hello")
}

// value type
func sayMessage(message string) {
	fmt.Println("sayMessage : ", message)
}

// reference type
func sayMessageWithPointer(message *string) {
	fmt.Println(*message)
	*message = "Golang"
}
