package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(lineEncoding("abbcabb"))
}

func lineEncoding(s string) string {
	output, counter := "", 1
	for i := 0; i <= len(s)-1; i++ {
		if i < len(s)-1 && s[i] == s[i+1] {
			counter++
			continue
		}
		if counter == 1 {
			output += string(s[i])
			continue
		}
		output += strconv.Itoa(counter) + string(s[i])
		counter = 1
	}
	return output
}
