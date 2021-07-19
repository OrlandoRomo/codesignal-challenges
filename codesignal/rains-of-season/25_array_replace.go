package main

func main() {

}

func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] == elemToReplace {
			inputArray[i] = substitutionElem
		}
	}
	return inputArray
}
