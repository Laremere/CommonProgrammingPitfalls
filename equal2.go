package main

import (
	"fmt"
)

func main() {
	var a, b int

	a = 1
	b = a
	a = 2

	fmt.Println("Variable A is ", a)
	fmt.Println("Variable B is ", b)
}
