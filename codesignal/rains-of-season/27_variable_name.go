package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("var_1__Int", variableName("var_1__Int"))
	fmt.Println("qq-q", variableName("qq-q"))
	fmt.Println("2w2", variableName("2w2"))
	fmt.Println("va[riable0", variableName("va[riable0"))
	fmt.Println("a a_2", variableName("a a_2"))
}
func variableName(name string) bool {
	regx := regexp.MustCompile("^\\d|\\s|[^a-zA-Z0-9_]")
	return !regx.Match([]byte(name))
}
