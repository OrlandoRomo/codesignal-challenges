package main

import "fmt"

func main() {
	fmt.Println(growingPlant(100, 10, 910))
	fmt.Println(growingPlant(10, 9, 4))
	fmt.Println(growingPlant(5, 2, 7))
	fmt.Println(growingPlant(10, 9, 4))
}

func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
	seed, nights := upSpeed, 0
	for seed < desiredHeight {
		seed = seed - downSpeed + upSpeed
		nights++
	}
	nights++
	return nights
}
