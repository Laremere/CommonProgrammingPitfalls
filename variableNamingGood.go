package main

import "fmt"

func main() {
	fmt.Println(SumConcecutiveNumbers(2, 2))
}

func SumConcecutiveNumbers(startingNumber int, numberToSum int) (sum int) {
	sum = 0
	currentNumber := startingNumber
	for i := 0; i < numberToSum; i++ {
		sum += numberToSum
		currentNumber += 1
	}
	return sum
}
