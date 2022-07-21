package main

func main() {

}

func countingValleys(steps int32, path string) int32 {
	runes := []rune(path)
	upCounter, downCounter, max, valleys := 0, 0, 0, 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == 'U' {
			if downCounter != 0 || max >= downCounter {
				max = downCounter
				valleys++
			}

			downCounter = 0
			upCounter++
		}
		if runes[i] == 'D' {
			upCounter = 0
			downCounter++
		}
	}
	return int32(valleys)
}
