package main

import (
	"reflect"
	"sort"
)

func main() {

}

func threeSum(nums []int) [][]int {
	output := make([][]int, 0)
	tmp := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			for k := 2; k < len(nums); k++ {
				if i != j && i != k && j != k && nums[i]+nums[j]+nums[k] == 0 {
					slice := []int{nums[i], nums[j], nums[k]}
					sort.Ints(slice)
					if reflect.DeepEqual(tmp, slice) {
						break
					}

					output = append(output, slice)
					tmp = slice
				}
			}
		}
	}

	return output
}
