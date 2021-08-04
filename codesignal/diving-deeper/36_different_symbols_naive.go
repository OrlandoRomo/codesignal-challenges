package main

import "fmt"

func main() {
	fmt.Println(differentSymbolsNaive("cabca"))
}

func differentSymbolsNaive(s string) int {
	chars := make(map[string]struct{}, 0)
	counter := 0
	for _, char := range s {
		if _, ok := chars[string(char)]; !ok {
			chars[string(char)] = struct{}{}
			counter++
		}
	}
	return counter
}
