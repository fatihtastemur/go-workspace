/*
Variadic Functions : Parametre sayısı bilinmeyen fonksiyonlar
*/
package main

import "fmt"

func main() {
	total1 := sumAllV1(7, 10, 58, 90, 82, 55)
	fmt.Println("Total 1:", total1)

	total2 := sumAllV2(7, 10, 58, 90, 82, 55)
	fmt.Println("Total 2:", total2)

	fmt.Println(concatenate("-", "foo", "bar", "baz"))
}

func sumAllV1(numbers ...int) int {
	total := 0
	for i := range numbers {
		total += numbers[i]
	}
	return total
}

func sumAllV2(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func concatenate(seperator string, strs ...string) string {
	result := ""
	for i, str := range strs {
		if i > 0 {
			result += seperator
		}
		result += str
	}
	return result
}
