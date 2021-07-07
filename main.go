package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 2}
	b := []int{2, 1, 1}
	fmt.Println(areSimilar(a, b))
}

func areSimilar(a []int, b []int) bool {
	m := make(map[int]int, 0)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		if _, ok := m[v]; ok {
			m[v]--
			if m[v] == 0 {
				delete(m, v)
			}
		}
	}
	return len(m) == 0
}
