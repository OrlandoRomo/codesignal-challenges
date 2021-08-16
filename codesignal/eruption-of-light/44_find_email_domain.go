package main

import "strings"

func main() {

}

func findEmailDomain(address string) string {
	domains := strings.Split(address, "@")
	return domains[len(domains)-1]
}
