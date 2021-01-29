package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//https://golang.org/pkg/fmt/#example_Sprintf
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}
