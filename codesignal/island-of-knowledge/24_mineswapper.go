package main

import (
	"fmt"
)

func main() {
	test := [][]bool{
		{false, false, false},
		{false, false, false},
	}
	fmt.Println(minesweeper(test))
}

func minesweeper(matrix [][]bool) [][]int {
	count := make([][]int, len(matrix))
	for i := range count {
		count[i] = make([]int, len(matrix[0]))
	}

	lenMatrix := len(matrix)
	lenSubMatrix := len(matrix[0])
	for i := 0; i < lenMatrix; i++ {
		for j := 0; j < lenSubMatrix; j++ {
			// right neighbor
			if j+1 < lenSubMatrix && matrix[i][j+1] == true {
				count[i][j]++
			}
			// down neighbor
			if i+1 < lenMatrix && matrix[i+1][j] == true {
				count[i][j]++
			}
			// left neighbor
			if j-1 > -1 && matrix[i][j-1] == true {
				count[i][j]++
			}
			// top neighbor
			if i-1 > -1 && matrix[i-1][j] == true {
				count[i][j]++
			}
			// top right diagonal
			if i-1 > -1 && j+1 < lenSubMatrix && matrix[i-1][j+1] == true {
				count[i][j]++
			}
			// top left diagonal
			if i-1 > -1 && j-1 > -1 && matrix[i-1][j-1] == true {
				count[i][j]++
			}
			// down right diagonal
			if i+1 < lenMatrix && j+1 < lenSubMatrix && matrix[i+1][j+1] == true {
				count[i][j]++
			}
			// down left diagonal
			if i+1 < lenMatrix && j-1 > -1 && matrix[i+1][j-1] == true {
				count[i][j]++
			}
		}
	}
	return count
}
