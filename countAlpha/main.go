package main

import "fmt"

func CountAlpha(s string) int {
	res := 0
	for _,val := range s {
		if (val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z') {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}