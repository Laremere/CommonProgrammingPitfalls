package main

import (
	"fmt"
)

const OneBillion = 1000000000

func main() {
	var fraction float32 = float32(1) / OneBillion
	var sum float32 = 0

	for i := 0; i < OneBillion; i++ {
		sum += fraction
	}

	fmt.Println(fraction * OneBillion)
	fmt.Println(sum)
}
