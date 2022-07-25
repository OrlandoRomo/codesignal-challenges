package main

func main(){

}

func KnapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int{
    if weight1+weight2 <= maxW {
        return value1 + value2
    }
    val:=0
    if weight1<= maxW {
        val = value1
    }
    if weight2<= maxW {
        if val < value2 {
            val = value2
        }
    }
    
    return val
}
