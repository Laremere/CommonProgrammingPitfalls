package main

import (
	"fmt"
)

func main() {
	var divisor int = 5
	var step float32 = 1 / float32(divisor)
	var total float32 = 0
	for i := 0; i < divisor; i++ {
		total += step
	}

	fmt.Println(total)
	fmt.Println(step * float32(divisor))
}
