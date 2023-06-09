package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(a uint) uint {
	if a == 0 {
		return 1
	}

	return a * factorial(a-1)
}
