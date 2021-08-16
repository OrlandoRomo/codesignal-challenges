package main

func main() {

}

func bishopAndPawn(bishop string, pawn string) bool {
	bishopCoord := []rune(bishop)
	pawnCoord := []rune(pawn)
	sum1 := bishopCoord[0] + bishopCoord[1]
	sum2 := pawnCoord[0] + pawnCoord[1]
	return sum1%2 == 0 && sum2%2 == 0 || sum1 == sum2
}
