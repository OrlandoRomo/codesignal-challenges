package main

func main() {}

func MetroCard(lastNumberOfDays int) []int {

	if lastNumberOfDays == 28 {
		return []int{31}
	}
	if lastNumberOfDays == 30 {
		return []int{31}
	}

	return []int{28, 30, 31}
}
