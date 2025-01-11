package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	res := make([]string, 0)
	word := ""
	for i,v := range s {
		if v != ' ' && v != '\t' && v != '\n' {
			word = word + string(v)
			if i == len(s)-1 {
				res = append(res, word)
			}
		} else {
			res = append(res, word)
			word = ""
		}
	}
	return res
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}