package main

func main() {

}

func avoidObstacles(inputArray []int) int {
	sum := 1
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i]%sum == 0 {
			i = -1
			sum++
		}
	}
	return sum
}
