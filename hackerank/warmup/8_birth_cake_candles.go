package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(birthdayCakeCandles([]int32{4, 4, 1, 3}))
}

func birthdayCakeCandles(candles []int32) int32 {
	size, counter, max := len(candles), 0, 0
	newArray := make([]int, 0)
	for _, v := range candles {
		newArray = append(newArray, int(v))
	}
	sort.Ints(newArray)
	max = newArray[size-1]
	for i := size - 1; i >= 0; i-- {
		if newArray[i] != max {
			break
		}
		counter++
	}

	return int32(counter)
}
