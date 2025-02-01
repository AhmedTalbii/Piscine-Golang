package main

import "fmt"

func RepeatAlpha(s string) string {
	res := ""
	for _,char := range s {
		if char >= 'a' && char <= 'z' {
			for i := 0 ; i < int(char-'a'+1) ; i++ {
				res = res + string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			for i := 0 ; i < int(char-'A'+1) ; i++ {
				res = res + string(char)
			}
		} else {
			res = res + string(char)
		}
	}
	return res
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}