package main

import "fmt"

func main() {
	fmt.Println(spiralNumbers(4))
}

func spiralNumbers(n int) [][]int {
	spiral, clock := make([][]int, n), 1
	for i := range spiral {
		spiral[i] = make([]int, n)
	}
	top, down, left, right := 0, n-1, 0, n-1
	direction := 0
	for top <= down && left <= right {
		if direction == 0 {
			for i := left; i <= right; i++ {
				spiral[top][i] = clock
				clock++
			}
			direction = 1
			top += 1
		}
		if direction == 1 {
			for i := top; i <= down; i++ {
				spiral[i][right] = clock
				clock++
			}
			direction = 2
			right -= 1
		}
		if direction == 2 {
			for i := right; i >= left; i-- {
				spiral[down][i] = clock
				clock++
			}
			direction = 3
			down -= 1
		}
		if direction == 3 {
			for i := down; i >= top; i-- {
				spiral[i][left] = clock
				clock++
			}
			left += 1
		}
		direction = (direction + 1) % n
	}
	return spiral
}
