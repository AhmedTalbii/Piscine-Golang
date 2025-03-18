package main

import "fmt"
func RecursiveFactorial(nb int) int {
	if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 1
}


func main() {
	age:= 4
	age = RecursiveFactorial(age)
	fmt.Println(age)
}
