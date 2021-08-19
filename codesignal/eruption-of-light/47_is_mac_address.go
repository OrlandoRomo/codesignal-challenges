package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(isMAC48Address("00-1B-63-84-45-E6"))
	fmt.Println(isMAC48Address("Z1-1B-63-84-45-E6"))
	fmt.Println(isMAC48Address("not a MAC-48 address"))
}

func isMAC48Address(inputString string) bool {
	_, err := net.ParseMAC(inputString)
	return err == nil
}
