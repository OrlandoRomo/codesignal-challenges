package main

func main() {

}

func solution(n int) int {
	// 29 = 29 / 10 = 29 % 10 = 9
	first, second := n/10, n%10
	return first + second
}
