package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(max - min)
	fmt.Println("Generate Number : ", number)
}
