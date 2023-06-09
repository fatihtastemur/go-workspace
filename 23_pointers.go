package main

import "fmt"

func main() {
	a := 8
	fmt.Println("a value is :")
	fmt.Println(a) //Output: 8

	fmt.Println("a pointer is :")
	fmt.Println(&a) //Output: 0xc000128008

	b := &a
	fmt.Println(b)  //Output: 0xc000128008
	fmt.Println(*b) //Output: 8

	x := 5
	y := 10
	fmt.Println("swap")
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
