package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))

}

//	func pivotIndex(nums []int) int {
//		sumLeft, sumRight, pivot := 0, 0, 0
//		for {
//			if pivot == len(nums) {
//				return -1
//			}
//			left := nums[:pivot]
//			rigt := nums[pivot+1:]
//			sumLeft, sumRight = sum(left...), sum(rigt...)
//			if sumLeft == sumRight {
//				return pivot
//			}
//			pivot++
//			sumLeft, sumRight = 0, 0
//		}
//
//		return pivot
//	}
func pivotIndex(nums []int) int {
	total := sum(nums...)
	leftSum := 0
	for i, val := range nums {
		rightSum := total - val - leftSum
		if rightSum == leftSum {
			return i
		}
		leftSum += val
	}

	return -1
}

func sum(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	sum := 0
	for _, val := range values {
		sum += val
	}

	return sum
}
