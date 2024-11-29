package main

import "fmt"

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}