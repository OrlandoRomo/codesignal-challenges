package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(kangaroo(43, 2, 70, 2))
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	x := x1 - x2
	v := v2 - v1
	if v == 0 {
		return "NO"
	}
	jumps := int32(math.Abs(float64(x / v)))
	if x1+(jumps*v1) == x2+(jumps*v2) {
		return "YES"
	}
	return "NO"
}
