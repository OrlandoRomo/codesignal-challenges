package main

import "fmt"

// for debuggin purposes
func main() {
	test := []int{10, 1, 2, 3, 4, 5}
	fmt.Println(almostIncreasingSequence(test))
}

func almostIncreasingSequence(sequence []int) bool {
	max := sequence[0]
	maxIndex := 0
	for i, num := range sequence {
		if max < num {
			max = num
			maxIndex = i
		}
	}
	for i := 0; i < len(sequence)-1; i++ {
		if i == maxIndex {
			continue
		}
		if i+1 == maxIndex {
			if sequence[i+1] >= sequence[i+2] {
				return false
			}
		}
		if sequence[i] >= sequence[i+1] {
			return false
		}
	}
	return true
}
