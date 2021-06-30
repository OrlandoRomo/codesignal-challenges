package main

func main() {

}

func allLongestStrings(inputArray []string) []string {
	max := len(inputArray[0])
	strs := make([]string, 0)
	for _, str := range inputArray {
		strLen := len(str)
		if strLen > max {
			max = strLen
			strs = nil
			strs = append(strs, str)
			continue
		}
		if strLen == max {
			strs = append(strs, str)
		}
	}
	return strs
}
