package main

import "fmt"

func Atoi(s string) int {
	result := 0
	isNegative := false
	start := 0

	
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		if s[0] == '-' {
			isNegative = true
		}
		start = 1 
	}

	for i := start; i < len(s); i++ {
		char := s[i]
		if char < '0' || char > '9' {
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}

	if isNegative {
		result = -result
	}

	return result
}

func BasicAtoi(s string) int {
	result := 0
	for _, char := range s {
		
		// fmt.Println("char :",char)
		digit := int(char - '0')
		// fmt.Println("digit :",digit)
		result = result*10 + digit
		// fmt.Println("result :",result)
	}
	return result
}

func BasicAtoi2(s string) int {
	result := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			// If not a digit, return 0
			return 0
		}
		// fmt.Println("char :",char)
		digit := int(char - '0')
		// fmt.Println("digit :",digit)
		result = result*10 + digit
		// fmt.Println("result :",result)
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi2("12345"))       
	fmt.Println(BasicAtoi2("0000000012345")) 
	// fmt.Println(BasicAtoi("000000"))        
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}