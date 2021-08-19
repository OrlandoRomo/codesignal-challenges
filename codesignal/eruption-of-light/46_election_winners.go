package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(electionsWinners([]int{2, 3, 5, 2}, 3))
	fmt.Println(electionsWinners([]int{5, 1, 3, 4, 1}, 0))
	// fmt.Println(electionsWinners([]int{1, 3, 3, 1}, 0))
	// fmt.Println(electionsWinners([]int{2,3,5,2}, 3))
}

func electionsWinners(votes []int, k int) int {
	winners := 0
	sort.Ints(votes)
	last, max := votes[len(votes)-1], 0
	for i := range votes {
		if votes[i] == last {
			max++
		}
		if votes[i]+k > last {
			winners++
		}
	}
	if k == 0 && max == 1 {
		return 1
	}
	return winners
}
