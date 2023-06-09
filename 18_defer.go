/*
Defer : Erteleme
*/
package main

import "fmt"

func main() {
	defer fmt.Println("First Sentence")
	fmt.Println("Second Sentence")
}

/*
	Output :
		Second Sentence
		First Sentence
*/
