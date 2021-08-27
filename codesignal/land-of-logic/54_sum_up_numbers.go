package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(sumUpNumbers("2 apples, 12 oranges"))
}

func sumUpNumbers(inputString string) int {
	regex := regexp.MustCompile("[0-9]+")
	numbers := regex.FindAllString(inputString, -1)
	sum := 0
	for _, num := range numbers {
		n, _ := strconv.Atoi(num)
		sum += n
	}
	return sum
}
