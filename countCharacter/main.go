package main 

import "fmt"

func CountChar(str string, c rune) int {
	res := 0 
    for _,val := range str {
		if val == c {
			res++
		}
    }
	return res
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}