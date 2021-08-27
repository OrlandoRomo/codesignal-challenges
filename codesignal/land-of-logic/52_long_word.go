package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(longestWord("Ready, steady, go!"))
}

func longestWord(text string) string {
	regex := regexp.MustCompile(`[^a-zA-Z ]+`)
	strs := strings.Split(regex.ReplaceAllString(text, ""), " ")
	maxStr := ""
	//O(N)
	for _, str := range strs {
		if len(str) > len(maxStr) {
			maxStr = str
		}
	}
	return maxStr
}
