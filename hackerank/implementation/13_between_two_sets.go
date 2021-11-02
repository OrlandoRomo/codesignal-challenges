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
	r := make([]int32, 0, 10)
	for i := a[len(a)-1]; i <= b[0]; i++ {
		j := 0
		for ; j < len(a); j++ {
			if i%a[j] != 0 {
				break
			}
		}
		if j == len(a) {
			k := 0
			for ; k < len(b); k++ {
				if b[k]%i != 0 {
					break
				}
			}
			if k == len(b) {
				r = append(r, i)
			}
		}
	}
	return int32(len(r))
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
