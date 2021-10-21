package main

import "fmt"

func main() {

}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	countFruits(s, t, a, apples)
	countFruits(s, t, b, oranges)
}

func countFruits(s int32, t int32, tree int32, fruits []int32) {
	counter := 0
	for i := range fruits {
		if s <= tree+fruits[i] && tree+fruits[i] <= t {
			counter++
		}
	}
	fmt.Println(counter)
}
