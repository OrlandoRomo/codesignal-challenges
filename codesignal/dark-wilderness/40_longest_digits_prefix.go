package main

import "regexp"

func main() {
	println(longestDigitsPrefix("123aa1"))
}

func longestDigitsPrefix(inputString string) string {
	regex := regexp.MustCompile("^([0-9]+)")
	return regex.FindString(inputString)
}
