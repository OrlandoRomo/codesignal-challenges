package main

func main() {}

func arithmeticExpression(a int, b int, c int) bool {
	if a+b == c || a-b == c || a*b == c || (a/b == c && a%b == 0) {
		return true
	}
	return false
}
