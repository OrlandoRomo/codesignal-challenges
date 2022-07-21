package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {

	i := 0
	for i < len(nums) {
		if i == len(nums) {
			if nums[i] == nums[i+1] {
				nums = append(nums[:i], nums[i+1:]...)
				i = 0
				continue
			}
		}
		i++
	}
	return len(nums)
}

// 0 0 1 1
