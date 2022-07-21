package main

import "fmt"

var (
	symbols = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	occurencies = make(map[string]int)
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	number := 0

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && symbols[string(s[i])] < symbols[string(s[i+1])] {
			number -= symbols[string(s[i])]
			continue
		}
		number += symbols[string(s[i])]
	}

	return number
}
