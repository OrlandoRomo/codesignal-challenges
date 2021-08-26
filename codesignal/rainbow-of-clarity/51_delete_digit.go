package main

import (
	"strconv"
	"strings"
)

func main() {
	deleteDigit(152)
}

func deleteDigit(n int) int {
	max, number := 0, strconv.Itoa(n)
	digits := strings.Split(number, "")
	digitsCopy := make([]string, len(digits))
	copy(digitsCopy, digits)
	for i := range digits {
		num := strings.Join(append(digitsCopy[:i], digitsCopy[i+1:]...), "")
		n, _ := strconv.Atoi(num)
		if n > max {
			max = n
		}
		digitsCopy = make([]string, len(digits))
		copy(digitsCopy, digits)
	}
	return max
}
