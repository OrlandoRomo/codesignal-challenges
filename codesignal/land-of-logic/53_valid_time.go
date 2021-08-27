package main

import "regexp"

func main() {

}

func validTime(time string) bool {
	regexp := regexp.MustCompile("^(([0-1][0-9]|2[0-3]):[0-5][0-9])$")
	return len(regexp.FindString(time)) != 0
}
