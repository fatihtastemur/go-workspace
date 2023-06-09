package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// input string
	fmt.Println("Please enter username : ")
	var username string
	fmt.Scanln(&username)
	fmt.Println("Username : ", username)

	// input int
	fmt.Println("Please enter number : ")
	var number int
	fmt.Scanln(&number)
	fmt.Println(number * number)

	// input sentence
	fmt.Println("Please enter sentence : ")
	inputReader := bufio.NewReader(os.Stdin)
	sentence, _ := inputReader.ReadString('\n')
	fmt.Println(sentence)
}
