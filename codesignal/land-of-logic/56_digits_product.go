package main

import (
	"fmt"
)

func main() {
	fmt.Println(digitsProduct(12))
}

func digitsProduct(product int) int {
	if product == 0 {
		return 10
	}
	if product == 1 {
		return 1
	}

	digits := []int{}
	for i := 9; i > 1; i-- {
		for (product % i) == 0 {
			digits = append(digits, i)
			product /= i
		}
	}
	if product != 1 {
		return -1
	}
	total := 0
	for i := len(digits) - 1; i >= 0; i-- {
		total = 10*total + digits[i]
	}

	return total
}
