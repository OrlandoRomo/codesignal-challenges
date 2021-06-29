package main

import "fmt"

// for debuggin purposes
func main() {
	// test1 := [][]int{
	// 	{1, 0, 3},
	// 	{0, 2, 1},
	// 	{1, 2, 0},
	// }
	// test2 := [][]int{
	// 	{1, 1, 1, 1},
	// 	{0, 5, 0, 0},
	// 	{2, 1, 3, 10},
	// }
	test3 := [][]int{
		{1, 1, 1, 0},
		{0, 5, 0, 1},
		{2, 1, 3, 10},
	}
	// fmt.Println(matrixElementsSum(tes1))
	// fmt.Println(matrixElementsSum(test2))
	fmt.Println(matrixElementsSum(test3))
}

func matrixElementsSum(matrix [][]int) int {
	sumCost, sumCostPerColumn := 0, 0
	buildingLen, officeLen := len(matrix), len(matrix[0])
	i, j := buildingLen-1, officeLen-1
	for j >= 0 {
		currentOffice := matrix[i][j]
		if i == 0 {
			if currentOffice == 0 {
				sumCost -= sumCostPerColumn
			} else {
				sumCost += currentOffice
			}
			j--
			i = buildingLen - 1
			sumCostPerColumn = 0
			continue
		}
		officeUp := matrix[i-1][j]
		if officeUp != 0 {
			sumCost += currentOffice
			sumCostPerColumn += currentOffice
		}
		// jump up to the office level
		i--
	}
	return sumCost
}
