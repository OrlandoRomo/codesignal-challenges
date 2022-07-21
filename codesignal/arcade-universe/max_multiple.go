packge main

func main(){}

func maxMultiple(divisor int, bound int) int {
   // 10 / 3 = 3 * 3 = 9
   n:= bound / divisor * divisor
    for n%divisor != 0 || n> bound || n<0 {
        n++
    }
    return n
}
