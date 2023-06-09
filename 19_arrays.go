package main

import "fmt"

func main() {
	// Array Definition 1
	var teams [6]string
	teams[0] = "Manchester United"
	teams[1] = "Liverpool"
	teams[2] = "Manchester City"
	teams[3] = "Chelsea"
	teams[4] = "Arsenal"
	teams[5] = "Tottenham"

	fmt.Println(teams)
	fmt.Println(teams[1])
	fmt.Println(len(teams))

	// Array Definition 2
	teamNicknames := [6]string{"UNITED", "REDS", "CITY", "BLUES", "GUNNERS", "SPURS"}
	fmt.Println(teamNicknames)

	unKnownSize := [...]int{99, 7, 10}
	fmt.Println(len(unKnownSize))
	fmt.Println(cap(unKnownSize))
}
