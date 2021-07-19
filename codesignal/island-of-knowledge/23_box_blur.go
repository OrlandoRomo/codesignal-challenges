package main

func main() {
	test := [][]int{
		{36, 0, 18, 9, 9, 45, 27},
		{27, 0, 254, 9, 0, 63, 90},
		{81, 255, 72, 45, 18, 27, 0},
		{0, 0, 9, 81, 27, 18, 45},
		{45, 45, 227, 227, 90, 81, 72},
		{45, 18, 9, 255, 9, 18, 45},
		{27, 81, 36, 127, 255, 72},
	}
	boxBlur(test)
}

func boxBlur(image [][]int) [][]int {
	box, row, column := 9, len(image), len(image[0])
	sum, minRow, minColumn := 0, 0, 0
	output := make([][]int, row-2)
	for i := range output {
		output[i] = make([]int, column-2)
	}
	for i := 0; i < row-2; i++ {
		for j := 0; j < column-2; j++ {
			for k := i; minRow < 3; k++ {
				for l := j; minColumn < 3; l++ {
					sum += image[k][l]
					minColumn++
				}
				minColumn = 0
				minRow++
			}
			sum /= box
			output[i][j] = sum
			minRow, sum = 0, 0
		}
	}
	return output
}
