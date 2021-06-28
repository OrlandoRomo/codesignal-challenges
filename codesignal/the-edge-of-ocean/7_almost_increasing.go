package main

import (
	"fmt"
)

func main() {
	test := []int{1, 2, 3, 4, 3, 6}
	fmt.Println(almostIncreasingSequence(test))
}

func almostIncreasingSequence(sequence []int) bool {
	sequenceCopy := make([]int, len(sequence))
	copy(sequenceCopy, sequence)
	i := findWrongPair(sequence)
	if i == -1 {
		return true
	}
	if findWrongPair(append(sequence[:i], sequence[i+1:]...)) == -1 {
		return true
	}
	if findWrongPair(append(sequenceCopy[:i+1], sequenceCopy[i+2:]...)) == -1 {
		return true
	}
	return false
}

func findWrongPair(sequence []int) int {
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			return i
		}
	}
	return -1
}
