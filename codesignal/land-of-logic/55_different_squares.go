package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 1},
		{2, 5, 6},
		{5, 7, 9},
		{1, 2, 3},
		{2, 2, 1},
	}
	fmt.Println(differentSquares(matrix))
}

func differentSquares(matrix [][]int) int {
	records := make(map[string]int)
	for x := 0; x < len(matrix)-1; x++ {
		for y := 0; y < len(matrix[x])-1; y++ {
			record := fmt.Sprintf("%v%v%v%v", matrix[x][y], matrix[x+1][y], matrix[x][y+1], matrix[x+1][y+1])
			records[record]++
		}
	}
	return len(records)
}
