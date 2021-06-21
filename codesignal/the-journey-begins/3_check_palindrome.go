package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkPalindrome("anita lava la tina"))
}

func checkPalindrome(inputString string) bool {
	newStr := strings.Replace(inputString, " ", "", -1)
	characters := []rune(newStr)
	for i, j := 0, len(characters)-1; i < j; i, j = i+1, j-1 {
		characters[i], characters[j] = characters[j], characters[i]
	}
	return newStr == string(characters)
}
