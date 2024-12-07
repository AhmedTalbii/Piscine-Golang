package main

import "fmt"

func IsUpper(s string) bool {
	res:= true
	for _,num := range s {
		if !(num >= 'A' && num <='Z') {
			res = false 
		}
	}
	return res
}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}