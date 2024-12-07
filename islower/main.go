package main 

import "fmt"

func IsLower(s string) bool {
	res := true 
	for _,num := range s {
		if !(num >= 'a' && num <= 'z' ) {
			res = false 
		}
	}
	return res
}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}