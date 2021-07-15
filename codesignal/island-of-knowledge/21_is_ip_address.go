package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(isIPv4Address("01.233.161.131"))
}

func isIPv4Address(inputString string) bool {
	ip := net.ParseIP(inputString)
	if ip == nil {
		return false
	}
	return ip.String() == inputString
}
