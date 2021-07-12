package main

func main() {

}

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	return (yourLeft == friendsLeft || yourLeft == friendsRight) && (yourRight == friendsLeft || yourRight == friendsRight)
}
