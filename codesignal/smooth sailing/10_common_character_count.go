package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonCharacterCount("", "adcaa"))
}

// Refactor this, did not like it :P
func commonCharacterCount(s1 string, s2 string) int {
	mapS1, mapS2 := countForMap(s1), countForMap(s2)
	counter := 0
	for letter, count := range mapS1 {
		if ctr, ok := mapS2[letter]; ok {
			if count > ctr {
				counter += ctr
				continue
			}
			counter += count
		}
	}
	return counter
}

func countForMap(str string) map[rune]int {
	countMap := make(map[rune]int, 0)
	for _, str := range str {
		countMap[str]++
	}
	return countMap
}
