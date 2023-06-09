package main

import "fmt"

func main() {
	teams := [6]string{"Manchester United", "Liverpool", "Manchester City", "Chelsea", "Arsenal", "Tottenham"}
	var group1 []string = teams[2:4] //slice

	fmt.Println(group1)
	fmt.Println(len(group1)) // len
	fmt.Println(cap(group1)) // capacity

	group2 := make([]string, 0, 5)            // make slice
	group2 = append(group2, "Leicester City") // append item
	group2 = append(group2, "Everton")
	fmt.Println(group2)

	// empty slice == nil
	var group3 []string
	if group3 == nil {
		fmt.Println("empty slice")
	}
}
