package main

import "fmt"

func main() {
	var team1, team2 = "SSC Napoli", "Atalanta BC"
	fmt.Println("Match of the day", team1, "-", team2)

	var team1Goal, team2Goal = 2, 2
	fmt.Printf("Match Scores : %d - %d\n", team1Goal, team2Goal)

	var ballPossession float32 = 50.7
	fmt.Printf("Home Ball Possesion : %g\n", ballPossession)
}
