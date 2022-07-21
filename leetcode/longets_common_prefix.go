package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	//	fmt.Println(longestCommonPrefix([]string{""}))
	fmt.Println(longestCommonPrefix([]string{"c", "acc", "ccc"}))

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
