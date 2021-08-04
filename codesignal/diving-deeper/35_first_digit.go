package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(firstDigit("var_1_das"))
	fmt.Println(firstDigit("q2q-q"))
	fmt.Println(firstDigit("0ss"))
}

func firstDigit(inputString string) string {
	regex := regexp.MustCompile("\\d")
	return regex.FindString(inputString)
}
