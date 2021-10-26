package main

func main() {

}

func breakingRecords(scores []int32) []int32 {
	res := make([]int32, 2)
	min, max := scores[0], scores[0]

	for i := 1; i < len(scores); i++ {
		if scores[i] < min {
			res[1]++
			min = scores[i]
		} else if scores[i] > max {
			res[0]++
			max = scores[i]
		}
	}
	return res
}
