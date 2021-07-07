package main

import "fmt"

func main() {
	fmt.Println(alternatingSums([]int{50, 60, 60, 45, 70}))
}

func alternatingSums(a []int) []int {
	result := make([]int, 2)
	for i, v := range a {
		fmt.Println(i % 2)
		result[i%2] += v
	}
	return result
}
