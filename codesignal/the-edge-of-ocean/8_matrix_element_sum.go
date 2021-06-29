package main

import "fmt"

func main() {
	test := [][]int{
		{1, 1, 1, 0},
		{0, 5, 0, 0},
		{2, 1, 3, 10},
	}
	fmt.Println(matrixElementsSum(test))
}

func matrixElementsSum(matrix [][]int) int {
	sumCost := 0
	buildingLen, officeLen := len(matrix), len(matrix[0])
	validRooms := make(map[int]bool)
	i, j := buildingLen, officeLen
	for i > buildingLen && j > officeLen {
		for ; j > officeLen; j-- {
			office := matrix[i][j]
			officeUp := matrix[i-1][j]
			if officeUp == 0 && office != 0 {
				// jump to the up office
				i--
				continue
			}

		}
	}
	return sumCost
}
