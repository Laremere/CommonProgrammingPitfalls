package main

import "fmt"

func main() {
	fmt.Println(someFunction(2, 2))
}

func someFunction(variable int, variable2 int) int {
	変数 := 0
	for variableB := 0; variableB < variable2; variableB++ {
		変数 += variable2
		variable += 1
	}
	return 変数
}
