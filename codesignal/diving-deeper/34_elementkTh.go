package main

import "fmt"

func main() {
	fmt.Println(extractEachKth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
}

func extractEachKth(inputArray []int, k int) []int {
	elements := make([]int, 0)
	n := 1
	for i, val := range inputArray {
		if i == n*k-1 {
			n++
			continue
		}
		elements = append(elements, val)
	}
	return elements
}
