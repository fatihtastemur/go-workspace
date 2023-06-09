/*
Closure Functions : Atanmış fonksiyonlar, değişkenlerin fonksiyon olarak tanımlanması
*/
package main

import "fmt"

func main() {
	total := func(x, y int) int {
		return x + y
	}

	fmt.Println("total :", total(2, 3))

	increment := outer()
	fmt.Println("increment : ", increment())
	fmt.Println("increment : ", increment())
	fmt.Println("increment : ", increment())
}

func outer() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
