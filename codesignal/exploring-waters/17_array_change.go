package main

import (
	"fmt"
)

func main() {
	test := []int{-1000, 0, -2, 0}
	fmt.Println(arrayChange(test))
}

func arrayChange(inputArray []int) int {
	counter, i := 0, 0
	for i < len(inputArray)-1 {
		if inputArray[i] == inputArray[i+1] || inputArray[i] > inputArray[i+1] {
			inputArray[i+1]++
			counter++
			continue
		}
		i++
	}
	return counter
}
