package main

import "fmt"

func Atoi(s string) int {
	start := 0
	result := 0
	negative := false
	if s[0] == '+' || s[0] == '-' {
		start = 1
		if s[0] == '-' {
			negative = true
		}
	}
	for i := start ; i < len(s) ; i++ {
		if s[i] <'0' || s[i] >'9' {
			return 0
		}
		digit := int(s[i] - '0')
		result = result*10 + digit
	}
	if negative {
		return -result
	}

	return result
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
	/*$ go run .
	12345
	12345
	0
	0
	1234
	-1234
	0
	0
	$*/
}