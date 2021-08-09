package main

import (
	"fmt"
)

func main() {
	// fmt.Println(digitDegree(5))
	fmt.Println(digitDegree(91))
	// fmt.Println(digitDegree(100))
	// fmt.Println(digitDegree(99))
	// fmt.Println(digitDegree(911))
}

func digitDegree(n int) int {
	degree, digits := 0, 0
	for n >= 10 {
		digits = 0
		for n != 0 {
			digits += n % 10
			n = n / 10
		}
		degree++
		n = digits
	}
	return degree
}
