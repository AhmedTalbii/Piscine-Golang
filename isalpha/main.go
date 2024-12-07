package main 

import "fmt"

func IsAlpha(s string) bool {
	var res bool
	res = true
	for _ , char := range s {
		if !(byte(char) >= 'a' && byte(char) <= 'z') &&
		   !(byte(char) >= 'A' && byte(char) <= 'Z') &&
		   !(byte(char) >= '0' && byte(char) <= '9') {
				res = false
		   }
	}
	return res
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}