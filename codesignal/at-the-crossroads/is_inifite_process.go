package main

func main() {}

func IsInfiniteProcess(a, b int) bool {
	return (b-a)%2 == 1 || a > b
}
