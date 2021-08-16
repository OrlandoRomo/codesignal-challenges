package main

import "strings"

func main() {

}

func findEmailDomain(address string) string {
	l := len(address)
	index := strings.Index(address, "@")
	address = address[index+1 : l]
	if strings.Count(address, "@") >= 1 {
		return findEmailDomain(address)
	}
	return address
}
