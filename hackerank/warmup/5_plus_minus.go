package main

import "fmt"

func main() {

}

const (
	Positive = iota
	Negative
	Zero
)

func plusMinus(arr []int32) {
	ratios := make([]int, 3)
	for _, v := range arr {
		if v > 0 {
			ratios[Positive]++
		}
		if v < 0 {
			ratios[Negative]++
		}
		if v == 0 {
			ratios[Zero]++
		}
	}
	for _, ratio := range ratios {
		fmt.Printf("%.6f", float64(ratio)/float64(len(arr)))
	}
}
