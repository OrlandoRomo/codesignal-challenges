package main

func main() {

}

func lastestNumber(n int) int {
	max := 1
	for i := 1; i <= n; i++ {
		max *= 10
	}
	return max - 1
}
