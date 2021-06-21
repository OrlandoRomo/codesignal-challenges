package main

import "fmt"

func main() {
	fmt.Println(shapeArea(3))
}

func shapeArea(n int) int {
	if n == 1 {
		return 1
	}
	area := 1
	for n-1 != 0 {
		area += (n - 1) * 4
		n--
	}
	return area
}
