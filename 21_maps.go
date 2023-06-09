package main

import "fmt"

type squad struct {
	goalkeeper, defender, midfielder, striker string
}

func main() {

	var colors map[string]string
	colors = make(map[string]string)
	colors["red"] = "#FF0000"
	colors["green"] = "#00FF00"
	colors["blue"] = "#0000FF"
	fmt.Println(colors)

	fmt.Println(colors["red"])

	// add map
	var mySquad = map[string]squad{
		"ace":     squad{"Muslera", "Nelsson", "Torreira", "Icardi"},
		"reserve": squad{"Jankat", "Ross", "Midtsjo", "Gomis"},
	}

	fmt.Println(mySquad["ace"])
	fmt.Println(mySquad["reserve"])
	fmt.Println(mySquad)

	// remove map
	order := make(map[string]int)
	order["no"] = 25
	fmt.Println(order["no"])
	delete(order, "no")
	fmt.Println(order["no"])
	fmt.Println(order)
}
