/*
	 	Data Types
		-----------------------------------------------------------------------------
		Basic Data Types     : Numbers, String, Boolean
		Cluster Data Types   : Array, Structure
		Reference Data Types : Pointers, slices, maps, functions, channels

		Numbers : int, float, complex
		---------------------------------------------------------------------------------
		------- int8   ---------------- -128   --------------  127       ----------------
		------- int16  ---------------- -32768 --------------  32767     ----------------
		------- int32  ---------------- -2^31  --------------  2^31 - 1  ----------------
		------- int64  ---------------- -2^63  --------------  2^63 - 1  ----------------
		------- uint8  ----------------  0     --------------  255       ----------------
		------- uint16 ----------------  0     --------------  65535     ----------------
		---------------------------------------------------------------------------------
*/
package main

import "fmt"

func main() {
	var minInt8 int8 = -128
	var maxInt8 int8 = 127

	fmt.Println("Min int8 value : ", minInt8)
	fmt.Println("Max int8 value : ", maxInt8)

	// int
	var x int = 99
	fmt.Println("Integer Number : ", x)

	// float
	var y float32 = 35.19
	fmt.Println("Float Number : ", y)

	// complex
	var z complex128 = complex(6, 2)
	fmt.Println("Complex Number : ", z)

	// string
	var str string = "Golang"
	fmt.Println("String")
	fmt.Println(str)
	fmt.Println("String : ", str)

	// boolean
	var isActive bool = false
	fmt.Println("Boolean : ", isActive)
}
