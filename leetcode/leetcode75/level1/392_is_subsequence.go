package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isSubsequence("ab", "baab"))
}

func isSubsequence(s string, t string) bool {
	currentIndex := 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(t, string(s[i]))
		if index == -1 {
			return false
		}

		if currentIndex > index {
			return false
		}
		t = strings.Repeat(" ", index+1) + t[index+1:]
		currentIndex = index
	}
	return true
}
