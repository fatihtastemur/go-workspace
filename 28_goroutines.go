package main

import (
	"fmt"
	"time"
)

func main() {
	go speak()
	time.Sleep(time.Second)
	fmt.Println("Goroutines")
}

func speak() {
	fmt.Println("Hello")
}
