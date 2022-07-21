package main

func main() {

}

func circleOfNumbers(n int, firstNumber int) int {
	half := n / 2
	if firstNumber+half > (n - 1) {
		return firstNumber - half
	}
	return firstNumber + half
}
