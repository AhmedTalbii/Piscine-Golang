package main 

import "fmt"

func IsNumeric(s string) bool {
	res := true 
	for _,num := range s {
		if !(byte(num) >= '0' && byte(num) <= '9') {
			res = false 
		}
	}
	return res
}

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}