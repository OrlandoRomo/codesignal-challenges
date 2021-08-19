func buildPalindrome(st string) string {
	i, letters := 0, ""
	if isPalindrome(st) {
		return st
	}
	for i < len(st)-1 {
		letters = string(st[i]) + letters
		if isPalindrome(st + letters) {
			return st + letters
		}
		i++
	}
	return st
}

func isPalindrome(s string) bool {
	characters := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		characters[i], characters[j] = characters[j], characters[i]
	}
	return s == string(characters)
}
