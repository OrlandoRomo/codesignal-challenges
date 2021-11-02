package main

import "fmt"

func main() {
	fmt.Println(birthday([]int32{
		4,
	}, 4, 1))
}

func birthday(s []int32, d int32, m int32) int32 {
	var counter, i, j, length int32 = 0, 0, 0, int32(len(s))
	for i = 0; i <= length-m; i++ {
		var sum int32
		for j = i; j < i+m; j++ {
			sum += s[j]
		}
		if sum == d {
			counter++
		}
	}
	return counter
}
