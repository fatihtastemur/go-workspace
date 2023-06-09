package main

import "fmt"

func main() {
	// for loop
	total1 := 0
	for i := 0; i < 10; i++ {
		total1 += i
	}

	fmt.Println("Total For : ", total1)

	// while loop
	total2 := 1
	for total2 < 9 {
		total2 += total2
		fmt.Println("Total While : ", total2)
	}
}
