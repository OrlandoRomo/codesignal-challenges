package main

import (
	"fmt"
)

func main() {
	fmt.Println(palindromeRearranging("aabb"))
}

func palindromeRearranging(inputString string) bool {
	letters := make(map[string]int, 0)
	for _, c := range inputString {
		if _, ok := letters[string(c)]; !ok {
			letters[string(c)] = 1
			continue
		}
		letters[string(c)]++
	}
	oddCount := 0
	for v := range letters {
		if letters[v]%2 != 0 {
			oddCount++
		}
	}
	if oddCount > 1 {
		return false
	}
	return true
}
