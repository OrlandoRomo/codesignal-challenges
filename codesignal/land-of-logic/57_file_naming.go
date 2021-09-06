package main

import (
	"fmt"
	"strconv"
)

func main() {
	example1 := []string{"a(1)",
		"a(6)",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a"}
	// example2 := []string{"doc",
	// 	"doc",
	// 	"image",
	// 	"doc(1)",
	// 	"doc"}

	fmt.Println(fileNaming(example1))
}

func fileNaming(names []string) []string {
	files := make(map[string]int)
	output := make([]string, 0)
	for _, name := range names {
		newName := name
		count, ok := files[newName]
		for ok {
			newName = name + "(" + strconv.Itoa(count) + ")"
			_, ok = files[newName]
			count++
		}
		files[newName]++

		output = append(output, newName)
	}
	return output
}
