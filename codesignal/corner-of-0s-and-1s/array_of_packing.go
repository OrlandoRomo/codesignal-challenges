package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrayOfPacking([]int{24, 85, 0}))
}

func arrayOfPacking(a []int) int {
	var value int
	for i, v := range a {
		value |= v << uint(i*8)
	}
	fmt.Println(2 | 3)
	return value
}
