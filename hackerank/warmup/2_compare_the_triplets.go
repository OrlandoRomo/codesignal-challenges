package main

const (
	Alice = iota
	Bob
)

func compareTriplets(a []int32, b []int32) []int32 {
	awarded := make([]int32, 2)
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			awarded[Alice]++
		}
		if a[i] < b[i] {
			awarded[Bob]++
		}
	}
	return awarded
}
