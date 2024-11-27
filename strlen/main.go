package main

import "fmt"

func StrLen(s string) int {
	////////// 1
	// o:=0
	// for i,_ := range s {
	// 	o=i+1
	// }
	////////// 2
	// o:= len(s)
	////////// 3 better one 
	o:=0
	for range s {
		o++
	}
	return o
	
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}