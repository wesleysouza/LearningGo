package main

import (
	"fmt"
)

//var declaration
var a = "Golang"
var z int //declaration without value
//visibility: package (package level scope)

//create types
type number int // subjacent type int
var n number = 10

func main() {
	numBytes, numErr := fmt.Println("Hello world", "in GO") //The function print return write number of bytes and errors
	fmt.Println(numBytes, numErr)

	// := gopher
	// declares and assigns inside code block
	// visibility: inside in block, after declare
	// The gopher do automatic typing (type of variable)

	// =
	// assigns value in variable exist

	//primitive types -> static
	b := 16        // numbers
	c := "strings" // strings
	d := true      //boolean

	fmt.Println(b, c, d)
	//output with printf
	fmt.Printf("x: %v, %T\n", b, c)

	//expressions -> evaluate and don't do an action
	e := 10 == 10 // statement (1 or more expressions)
	f := 5 + 8
	fmt.Println(e)
	fmt.Println(f)

	test(10)

	fmt.Println(z)

	//Strings (interpreted string literals vs. raw string literals)

	str1 := "Bom dia\n Tudo bem?"
	str2 := `Bom dia\n Tudo bem?`

	fmt.Println(str1)
	fmt.Println(str2)

	//convertion of types
	//using type number
	numberInt := int(n)
	fmt.Println(numberInt)
}

func test(x int) {
	fmt.Println(x)
}
