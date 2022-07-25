package main

func main() {}

func ExtraNumber(a int, b int, c int) int {
	if a == b {
		return c
	}

	if a == c {
		return b
	}

	return a
}
