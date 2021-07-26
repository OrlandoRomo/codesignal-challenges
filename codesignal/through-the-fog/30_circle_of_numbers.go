package main

import (
	"fmt"
)

func main() {
	fmt.Println(circleOfNumbers(10, 2))
	fmt.Println(circleOfNumbers(10, 7))
	fmt.Println(circleOfNumbers(4, 1))
	fmt.Println(circleOfNumbers(6, 3))
}
func circleOfNumbers(n int, firstNumber int) int {
	return (firstNumber + n/2) % n
}
