package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isLucky(123321))
}

func isLucky(n int) bool {
	str := strconv.Itoa(n)
	strLength := len(str)
	firstHalf := str[:strLength/2]
	secondHalf := str[strLength/2:]
	sum1, sum2 := 0, 0
	for i := 0; i < strLength/2; i++ {
		n, _ := strconv.Atoi(string(firstHalf[i]))
		sum1 += n
		n, _ = strconv.Atoi(string(secondHalf[i]))
		sum2 += n
	}
	return sum1 == sum2
}
