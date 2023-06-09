package main

import "fmt"

type player struct {
	name     string
	number   int
	position string
}

func main() {
	player1 := player{"Icardi", 9, "F"}
	fmt.Println(player1)

	player2 := player{}
	player2.name = "Muslera"
	player2.number = 1
	player2.position = "GK"
	fmt.Println(player2)

	// Anonymous Struct
	car := struct {
		brand string
		model string
		year  int
	}{"Volvo", "XC40 T3 FWD Momentum", 2020}
	fmt.Println(car)
}
