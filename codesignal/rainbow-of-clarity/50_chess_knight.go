package main

func main() {

}

func chessKnight(cell string) int {
	total := 8
	cellX, cellY := cell[0], cell[1]-'0'

	if cellX == 'a' || cellX == 'h' {
		total = total / 2
	}
	if cellX == 'b' || cellX == 'g' {
		total = total - 2
	}
	if cellY == 1 || cellY == 8 {
		total = total / 2
	}
	if cellY == 2 || cellY == 7 {
		total = total - 2
	}
	return total
}
