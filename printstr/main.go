package main 

import "fmt"

func PrintStr(s string) {
	
	for _,val := range s {
	
		fmt.Printf("%c",val)
	}
	// r := string(arr)
	// fmt.Printf("%v",r)
}

func main() {
	PrintStr("Hello World!")
}