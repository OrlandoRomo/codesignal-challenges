package main

import "fmt"

func main() {
	test := []int{6, 2, 3, 8}
	fmt.Println(makeArrayConsecutive2(test))
}

func makeArrayConsecutive2(statues []int) int {
	min, max := statues[0], statues[0]
	for _, statue := range statues {
		if max < statue {
			max = statue
		}
		if min > statue {
			min = statue
		}
	}
	return (max - min + 1) - len(statues)
}
