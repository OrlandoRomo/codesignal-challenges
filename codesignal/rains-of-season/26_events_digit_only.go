package main

import (
	"strconv"
	"strings"
)

func main() {

}

func evenDigitsOnly(n int) bool {
	str := strconv.Itoa(n)
	digits := strings.Split(str, "")
	for _, digit := range digits {
		if n, err := strconv.Atoi(digit); err == nil && n%2 != 0 {
			return false
		}
	}
	return true
}
