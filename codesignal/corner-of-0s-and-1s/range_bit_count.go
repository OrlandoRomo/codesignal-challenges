package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(rangeBitCount(2, 7))
}

func rangeBitCount(a, b int) int {
	var count int
	for i := a; i <= b; i++ {
		count += bits.OnesCount(uint(i))
	}
	return count
}

//func rangeBitCount(a int, b int) int {
//	if a < 0 || b < a {
//		return 0
//	}
//	var count int
//	for _, num := range generator(a, b) {
//		binary := strconv.FormatInt(int64(num), 2)
//		count += strings.Count(binary, "1")
//	}
//
//	return count
//}
//
//func generator(a, b int) []int {
//	numbers := make([]int, 0)
//	for i := a; i <= b; i++ {
//		numbers = append(numbers, i)
//	}
//
//	return numbers
//}
