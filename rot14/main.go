package main 

import "github.com/01-edu/z01"

func Rot14(str string) string {
	result := ""
	for _, res:= range str{
		if rune(res) >= 'A' && rune(res) < 'M' || rune(res) >= 'a' && rune(res) < 'm' {
			result = result + string(rune(res) + 14)
		} else if rune(res) >= 'M' && rune(res) <= 'Z' || rune(res) >= 'm' && rune(res) <= 'z' {
			result = result + string(rune(res) - 12)
		} else {
			result = result + string(res)
		}
	}
	return result
}

func main() {
	result := Rot14("Hello! How are You?")
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}