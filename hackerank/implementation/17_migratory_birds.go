package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
}

func migratoryBirds(arr []int32) int32 {
	birds := make(map[int32]int)
	highest := 0
	for i := 0; i < len(arr); i++ {
		birds[arr[i]]++
		birdCouting := birds[arr[i]]
		if highest < birdCouting {
			highest = birdCouting
		}
	}
	ids := make([]int32, 0)
	for key, value := range birds {
		if value == highest {
			ids = append(ids, key)
		}
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	return ids[0]
}
