package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("2017")
}

func dayOfProgrammer(year int32) string {
	date := time.Date(int(year), 1, 1, 0, 0, 0, 0, time.UTC)
	dateEightMonths := time.Date(int(year), 9, 1, 0, 0, 0, 0, time.UTC)
	if year%4 == 0 {
		diff := dateEightMonths.Sub(date)
		fmt.Println(diff.Hours() / 24)
	}
	return ""
}
