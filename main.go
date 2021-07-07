package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseInParentheses("(bar)"))
}

func reverseInParentheses(inputString string) string {
	chars := []rune(inputString)
	i, left, right, length := 0, -1, -1, len(chars)
	var stack []string
	for i < length {
		if isOpened(chars[i]) {
			left = i
			i++
			stack = nil
			continue
		}
		if left != -1 {
			if isClosed(chars[i]) {
				right = i
				reverse := reverseString(stack)
				inputString = strings.ReplaceAll(inputString, inputString[left:right+1], reverse)
				chars = []rune(inputString)
				i, left, right, length = 0, -1, -1, len(chars)
				continue
			}
			stack = append(stack, string(chars[i]))
		}
		i++
	}
	return string(chars)
}
func isOpened(r rune) bool {
	return r == 40
}

func isClosed(r rune) bool {
	return r == 41
}

func reverseString(s []string) string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, "")
}
