package main

import "time"

func main() {

}
func timeConversion(s string) string {
	// Write your code here
	parsed, err := time.Parse("03:04:05PM", s)
	if err != nil {
		return ""
	}
	return parsed.Format("15:04:05")
}
