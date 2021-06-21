package main

import "fmt"

func main() {
	fmt.Println(centuryFromYear(1999))
}

func centuryFromYear(year int) int {
	return (year + 99) / 100
}
