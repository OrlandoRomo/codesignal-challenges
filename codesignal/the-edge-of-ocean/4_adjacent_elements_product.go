package main

import "fmt"

func main() {
	test := []int{3, 6, -2, -5, 7, 3}
	fmt.Println(adjacentElementsProduct(test))
}

func adjacentElementsProduct(inputArray []int) int {
	max := inputArray[0] * inputArray[1]

	for i := 1; i < len(inputArray)-1; i++ {
		if max < inputArray[i]*inputArray[i+1] {
			max = inputArray[i] * inputArray[i+1]
		}
	}

	return max
}
