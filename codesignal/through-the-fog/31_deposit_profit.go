package main

func main() {

}

func depositProfit(deposit int, rate int, threshold int) int {
	float := float64(deposit)
	years := 0
	for int(float) < threshold {
		profit := float * float64(rate) / 100
		float += profit
		years++
	}
	return years
}
