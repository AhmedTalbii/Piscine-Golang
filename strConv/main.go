package main 

import (
	"fmt"
	"os"
)

func Atoi(str string) int {
	res := 0
	negative := false
	started := false
	for _, char := range str {
		if char == ' ' && !started {
			continue
		}
		if char == '-' && !started {
			negative = true
			started = true
		} else if char == '+' && !started {
			started = true
		} else if char >= '0' && char <= '9' {
			started = true
			res = res*10 + int(char-'0')
		} else {
			return 0
		}
	}
	if negative {
		return -res
	}
	return res
}

func applyFunc(f func(string) int,arr []string) int {
	if len(arr) != 3 {
		fmt.Println("you wrote an incorrect number of arguments. Example: 'number operator number'")
		 
	}

	num1 := f(arr[0])
	num2 := f(arr[2])

	if num1 == 0 && arr[0] != "0" {
		os.Exit(0)
	}
	if num2 == 0 && arr[2] != "0" {
		os.Exit(0)		
	}
	var res int

	switch arr[1] {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("No division by 0")
			return 0 
		}
		res = num1 / num2
	case "%":
		if num2 == 0 {
			fmt.Println("No modulo by 0")
			return 0
		}
		res = num1 % num2
	default:
		fmt.Println("Invalid operator")
		return 0
	}
	return res
}

func main() {
	res := applyFunc(Atoi, os.Args[1:])
	fmt.Println(res)
}