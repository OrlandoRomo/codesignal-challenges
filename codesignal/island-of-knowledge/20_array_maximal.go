package main

import "fmt"

func main() {
	test := []int{2, 4, 1, 0}
	fmt.Println(arrayMaximalAdjacentDifference(test))
}

func arrayMaximalAdjacentDifference(inputArray []int) int {
	max, i := -15, 0
	for i < len(inputArray)-1 {

		if inputArray[i]-inputArray[i+1] > max {
			max = inputArray[i] - inputArray[i+1]
		}
		if inputArray[i+1]-inputArray[i] > max {
			max = inputArray[i+1] - inputArray[i]
		}
		i++
	}
	return max
}
