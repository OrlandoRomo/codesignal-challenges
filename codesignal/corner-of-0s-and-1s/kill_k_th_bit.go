package main

import (
	"fmt"
)

func main() {
	fmt.Println(killKthBit(37, 3))
}

func killKthBit(n, k int) int {
	return n &^ (1 << uint(k-1))
}
