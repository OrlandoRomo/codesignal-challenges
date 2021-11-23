package main

import "fmt"

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, int32(1), int32(12))
}

func bonAppetit(bill []int32, k int32, b int32) {
	sum := 0
	pay := 0
	for i := 0; i < len(bill); i++ {
		if int32(i) == k {
			continue
		}
		sum += int(bill[i])
	}
	sum = sum / 2
	pay = int(b) - sum
	if pay == 0 {
		fmt.Println("Bon Appetit")
		return
	}
	fmt.Println(pay)
}
