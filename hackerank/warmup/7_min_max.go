package main

import (
	"fmt"
	"sort"
)

func main() {
	miniMaxSum([]int32{7, 69, 2, 221, 8974})
}

func miniMaxSum(arr []int32) {
	newArray := make([]int, 0)
	for _, v := range arr {
		newArray = append(newArray, int(v))
	}
	sort.Ints(newArray)
	min, max, l := 0, 0, len(newArray)
	for _, v := range newArray[:l-1] {
		min += v
	}

	for _, v := range newArray[1:] {
		max += v
	}
	fmt.Printf("%d %d\n", min, max)
}
