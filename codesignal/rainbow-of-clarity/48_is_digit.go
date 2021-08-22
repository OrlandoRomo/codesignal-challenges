package main

import "strconv"

func main() {

}

func isDigit(symbol string) bool {
	_, err := strconv.Atoi(symbol)
	return err == nil
}
