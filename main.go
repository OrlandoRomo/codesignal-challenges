package main

import "fmt"

// for debugging purposes
func main() {
	test1 := []string{"aba", "aa", "ad", "vcd", "aba", "abcd"}
	// test2 := [][]int{
	// 	{1, 1, 1, 1},
	// 	{0, 5, 0, 0},
	// 	{2, 1, 3, 10},
	// }
	// test3 := [][]int{
	// 	{1, 1, 1, 0},
	// 	{0, 5, 0, 1},
	// 	{2, 1, 3, 10},
	// }
	fmt.Println(allLongestStrings(test1))
	// fmt.Println(allLongestStrings(test2))
	// fmt.Println(allLongestStrings(test3))
}

func allLongestStrings(inputArray []string) []string {
	max := len(inputArray[0])
	strs := make([]string, 0)
	for _, str := range inputArray {
		strLen := len(str)
		if strLen > max {
			max = strLen
			strs = nil
			strs = append(strs, str)
			continue
		}
		if strLen == max {
			strs = append(strs, str)
		}
	}
	return strs
}
