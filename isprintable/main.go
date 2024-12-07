package main 

import "fmt"

func IsPrintable(s string) bool {
	res := true
	for _,char := range s {
		if char < 32 || char > 126{
			res = false
		}
	}
	return res
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}