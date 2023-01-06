package main

import (
	"fmt"
)

func main() {
	fmt.Println(isIsomorphic("bacd", "baba"))
}

func isIsomorphic(s string, t string) bool {
	mapST, mapTS := make(map[byte]byte, 0), make(map[byte]byte, 0)
	for i := 0; i < len(s); i++ {
		char1, ok1 := mapST[s[i]]
		char2, ok2 := mapTS[t[i]]

		if ok1 && char1 != t[i] {
			return false
		}

		if ok2 && char2 != s[i] {
			return false
		}

		mapST[s[i]] = t[i]
		mapTS[t[i]] = s[i]
	}

	return true
}

func mapString(s1, s2 string) map[string]string {
	chars := make(map[string]string, 0)
	for i := 0; i < len(s1); i++ {
		chars[string(s1[i])] = string(s2[i])
	}

	return chars
}
