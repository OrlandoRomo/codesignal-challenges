package main

func main() {

}
func sudoku(grid [][]int) bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			// iEl holds every submatrix row index
			iEl := i * 3
			// jEl hold every submatrix column index
			jEl := j * 3
			cellsChecked := make([]int, 9)
			for k := 0; k <= 2; k++ {
				for m := 0; m <= 2; m++ {
					// iBig holds the row position of the current element of th submatrix
					iBig := iEl + k
					// jBig holds the column position of the current element of th submatrix
					jBig := jEl + m
					val := grid[iBig][jBig]
					if val > 0 && val < 10 {
						if cellsChecked[val-1] == 1 {
							return false
						}
						cellsChecked[val-1] = 1
					}
				}
			}
		}
	}
	// Rows, Columns
	for i := 0; i <= 8; i++ {
		row, col := 0, 0
		for j := 0; j <= 8; j++ {
			row += grid[i][j]
			col += grid[j][i]
		}
		if row != 45 || col != 45 {
			return false
		}
	}
	return true
}
