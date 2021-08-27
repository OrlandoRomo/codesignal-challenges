package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(longestWord("Ready, steady, go!"))
}

func longestWord(text string) string {
	longest := ""
	regex := regexp.MustCompile("[A-Za-z]+")
	matches := regex.FindAllString(text, -1)
	for _, match := range matches {
		if len(match) > len(longest) {
			longest = match
		}
	}
	return longest
}
