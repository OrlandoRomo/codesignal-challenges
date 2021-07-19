package main

func main() {

}

func evenDigitsOnly(n int) bool {
	for n > 0 {
		digit := n % 2
		if digit%2 != 0 {
			return false
		}
		n = n / 10
	}
	return true
}
