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
	for i := 3 ; i*i <= nb ; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true 
}

func FindNextPrime(nb int) int {
	for i := nb ; i >= 0 ; i++{
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}