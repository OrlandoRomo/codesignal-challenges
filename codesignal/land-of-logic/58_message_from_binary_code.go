package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(messageFromBinaryCode("010010000110010101101100011011000110111100100001"))
}

func messageFromBinaryCode(code string) string {
	message, eightBits := "", 8
	for i := 0; i < len(code); i += eightBits {
		binary, _ := strconv.ParseInt(code[i:i+eightBits], 2, 64)
		message += string(rune(binary))
	}
	return message
}
