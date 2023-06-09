package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myString = "7"
	var myInt = 10

	// str to int
	toInt, _ := strconv.Atoi(myString)
	fmt.Println(toInt)

	// int to str
	toStr := strconv.Itoa(myInt)
	fmt.Println(toStr)

	// CASTING
	var i int = 89
	var f float64 = float64(i)
	var u uint = uint(f)
	var s string = string(i)

	fmt.Println(i, f, u, s)
}
