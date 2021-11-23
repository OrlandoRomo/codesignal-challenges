package main

import "fmt"

func main() {
	fmt.Println(sockMerchant(7, []int32{1, 2, 1, 2, 1, 3, 2}))
}

// Big O(1)
func sockMerchant(n int32, ar []int32) int32 {
	socks := make(map[int32]int32)
	var counter int32
	for _, color := range ar {
		socks[color]++
		if socks[color]%2 == 0 {
			counter++
		}
	}
	return counter
}
