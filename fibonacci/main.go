package main 

import "fmt"

func Fibonacci(index int) int {
    if index < 0 {
		return -1
	}
	f := 0 
	s := 1
	var res int
	for i := 2 ; i <= index ; i++ {
		res = s + f
		f = s
		s= res
		fmt.Println(res)
	}
	return res
}


// func Fibonacci(index int) int {
// 	if index < 0 {
// 		return -1
// 	}
// 	if index == 0 {
// 		return 0
// 	}
// 	if index == 1 {
// 		return 1
// 	} else {
// 		return Fibonacci(index-1) + Fibonacci(index-2)
// 	}
// }


func main() {
	arg1 := 2000
	fmt.Println(Fibonacci(arg1))
}