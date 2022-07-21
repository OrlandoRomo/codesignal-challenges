package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution("Yo soy tu padre BUAJAJA y se me olvidaba decirte que h"))
}

func solution(text string) string {
	chars := strings.Split(text, " ")
	for i, ch := range chars {
		chars[i] = encode(ch)
	}
	return strings.Join(chars, " ")
}

func encode(text string) string {
	if len(text) == 1 {
		return " "
	}

	if len(text) <= 3 {
		return text
	}

	chars, length := len(text)-2, len(text)

	return fmt.Sprintf("%v%v%v", string(text[0]), chars, string(text[length-1]))
}
