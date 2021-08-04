package main

func main() {

}

func arrayMaxConsecutiveSum(inputArray []int, k int) int {
	max := 0
	for i := 0; i < len(inputArray); i++ {
		sum := 0
		for j := i; j < k+i && j < len(inputArray); j++ {
			sum += inputArray[j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
