package main

import "fmt"

func main() {
	var average int = 75

	if average > 90 {
		fmt.Println("Student Point : A")
	} else if average > 80 {
		fmt.Println("Student Point : B")
	} else if average > 60 {
		fmt.Println("Student Point : C")
	} else {
		fmt.Println("Student Point : D")
	}
}
