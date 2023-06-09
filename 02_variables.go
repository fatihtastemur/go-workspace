package main

import "fmt"

// Global Variables
var globalVariable = "Global Variable"

var (
	firstName = "Fatih"
	lastName  = "TAŞTEMUR"
)

func main() {

	fmt.Println("Global Variable 1 : ", globalVariable)
	fmt.Println("Global Variable 2 : ", firstName)
	fmt.Println("Global Variable 3 : ", lastName)

	// Variable Definition
	var fullName string
	var age int
	var isActive bool
	var username = "ftastemur"

	fullName = "Aras Buğra TAŞTEMUR"
	age = 2
	isActive = true

	fmt.Println("Full Name : ", fullName)
	fmt.Println("Age : ", age)
	fmt.Println("Activate : ", isActive)
	fmt.Println("Username : ", username)

	/*
		Mathematical Operations
	*/
	fmt.Println(3 * 2)
	fmt.Println(12.5 / 4)
	fmt.Println(7 + 1)
	fmt.Println(21 - 2)

	// Constant
	const pi = 3.14
	fmt.Println("PI : ", pi)

	// Inferring Type
	minute := 60
	hour := 24
	fmt.Println("One day is ", minute*hour, "minutes.")

	// Integer
	var a1, b1 = 3, 5
	fmt.Println("a1, b1 : ", a1, b1)

	x, y, z := true, 8.5, "text"
	fmt.Println("x, y, z :", x, y, z)

	// Float
	var myFloat32 = 47.9
	fmt.Println(myFloat32)

	// Complex
	var myComplex = complex(4, 3)
	fmt.Println(myComplex)
}
