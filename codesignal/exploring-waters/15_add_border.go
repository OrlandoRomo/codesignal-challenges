package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(addBorder([]string{"abc", "ded"}))
}

func addBorder(picture []string) []string {
	starLength := len(picture[0]) + 2
	star := strings.Repeat("*", starLength)
	for i := 0; i < len(picture); i++ {
		picture[i] = string("*") + picture[i] + string("*")

	}
	picture = append([]string{star}, picture...)
	picture = append(picture, star)
	return picture
}
