package main

func main() {

}

func isBeautifulString(inputString string) bool {
	letters := [26]rune{}
	for _, v := range inputString {
		c := v - 'a'
		letters[c]++
	}

	for i := 1; i < len(letters); i++ {
		if letters[i] > letters[i-1] {
			return false
		}
	}
	return true
}
