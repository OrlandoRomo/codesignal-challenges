package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(chessBoardCellColor("A2", "A5"))
}

func getIndexIAndJ(cell1 string) (int, int) {
	var letters = "ABCDEFGH"
	cells := strings.Split(cell1, "")
	c1i, c2j := cells[0], cell1[1]
	i := strings.Index(letters, c1i)
	j, _ := strconv.Atoi(string(c2j))

	return i + 1, j
}
func chessBoardCellColor(cell1 string, cell2 string) bool {
	c1i, c1j := getIndexIAndJ(cell1)
	c2i, c2j := getIndexIAndJ(cell2)
	isBrownPosC1 := c1i == c1j || c1i%2 == 0 && c1j%2 == 0 || c1i%2 != 0 && c1j%2 != 0
	isBrownPosC2 := c2i == c2j || c2i%2 == 0 && c2j%2 == 0 || c2i%2 != 0 && c2j%2 != 0
	if isBrownPosC1 != isBrownPosC2 {
		return false
	}
	return true
}
