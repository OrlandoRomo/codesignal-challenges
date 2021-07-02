package main

import "fmt"

func main() {
	test := []int{-1, 150, 190, 170, -1, -1, 160, 180}
	fmt.Println(test)
	fmt.Println(sortByHeight(test))
}

func sortByHeight(a []int) []int {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] == -1 {
				break
			}
			if a[j] == -1 {
				continue
			}
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
