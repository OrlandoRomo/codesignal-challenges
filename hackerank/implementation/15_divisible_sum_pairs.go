package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(divisibleSumPairs(6, 5, []int32{1, 2, 3, 4, 5, 6}))
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	var counter, i, j int32
	for i = 0; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				counter++
			}
		}
	}
	return counter
}
