package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
