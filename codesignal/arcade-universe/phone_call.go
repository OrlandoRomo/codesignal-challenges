package main

func main() {}

func phoneCall(min1 int, min2_10 int, min11 int, S int) int {
	if min1 == S {
		return 1
	}
	if min1 > S {
		return 0
	}
	if min1+(9*min2_10) > S {
		return (S-min1)/min2_10 + 1
	}
	return ((S - (min1 + (9 * min2_10))) / min11) + 10
}
