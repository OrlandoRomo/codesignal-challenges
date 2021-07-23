package main

import "fmt"

func main() {
	fmt.Println(evenDigitsOnly(24687))
}

func evenDigitsOnly(n int) bool {
	for n > 0 {
		digit := n / 10
		n = n / 10
		if digit%2 != 0 {
			return false
		}
	}
	return true
}
