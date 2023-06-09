package main

import "fmt"

func main() {
	teams := [6]string{"Manchester United", "Liverpool", "Manchester City", "Chelsea", "Arsenal", "Tottenham"}

	for index, value := range teams {
		fmt.Printf("%d. index = %s\n", index, value)
	}
}
