package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 1, 3}
	fmt.Println(areSimilar(a, b))
}

func areSimilar(a []int, b []int) bool {
	notMath := 0
	for i := range a {
		if a[i] != b[i] {
			notMath++
		}
	}
	sort.Ints(a)
	sort.Ints(b)
	for i := range a {
		for a[i] != b[i] {
			return false
		}
	}
	return notMath < 3
}
