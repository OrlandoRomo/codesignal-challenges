package main

import "fmt"

func main() {
	fmt.Println(getTotalX([]int32{
		10, 10,
	}, []int32{
		100, 99, 98, 97, 96, 95, 94, 93, 92, 91,
	}))
}

func getTotalX(a []int32, b []int32) int32 {
	mcmA := getMCM(a, len(a))
	mcdB, counter := b[0], 0
	for i := 1; i < len(b); i++ {
		mcdB = getMCD(b[i], mcdB)
	}

	for i := mcmA; i <= int(mcdB); i += mcmA {
		if int(mcdB)%i == 0 {
			counter++
		}
	}
	return int32(counter)
}

// Get the mínimo común multiple using the Euclides algorithm
func getMCM(a []int32, size int) int {
	minimun := a[0]
	for i := 1; i < size; i++ {
		minimun = a[i] * minimun / getMCD(a[i], minimun)
	}
	return int(minimun)
}

// Get the máximo común divisor
func getMCD(a, b int32) int32 {
	if b == 0 {
		return a
	}
	return getMCD(b, a%b)
}
