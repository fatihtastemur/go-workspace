package main

import "fmt"

func doProcess(input int) (int, int) {
	process1 := input / 2
	process2 := input / 4

	return process1, process2
}

func main() {
	divideTo2, _ := doProcess(16)
	fmt.Println(divideTo2)
}
