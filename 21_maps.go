package main

import "fmt"

type squad struct {
	goalkeeper, defender, midfielder, striker string
}

func main() {
	// Definition
	var colors map[string]string
	colors = make(map[string]string)
	colors["red"] = "#FF0000"
	colors["green"] = "#00FF00"
	colors["blue"] = "#0000FF"
	fmt.Println(colors)
	fmt.Println(colors["red"])

	var mySquad = map[string]squad{
		"ace":     squad{"Muslera", "Nelsson", "Torreira", "Icardi"},
		"reserve": squad{"Jankat", "Ross", "Midtsjo", "Gomis"},
	}

	fmt.Println(mySquad["ace"])
	fmt.Println(mySquad["reserve"])
	fmt.Println(mySquad)

	// remove item from map
	states := make(map[string]string)
	states["IST"] = "İstanbul"
	states["ANK"] = "Ankara"
	states["IZM"] = "İzmir"

	fmt.Println(states)
	fmt.Println(states["ANK"])
	delete(states, "IZM")
	fmt.Println(states)

	// add item to map
	states["ANT"] = "Antalya"
	fmt.Println(states)
}
