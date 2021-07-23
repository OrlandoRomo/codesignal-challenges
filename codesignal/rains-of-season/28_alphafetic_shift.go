package main

import "fmt"

func main() {
	fmt.Println(alphabeticShift("crazy"))
}

func alphabeticShift(inputString string) string {
	charaters := []rune(inputString)
	for i := range charaters {
		if charaters[i] == 'z' {
			charaters[i] = 'a'
			continue
		}
		charaters[i]++
	}
	return string(charaters)
}
