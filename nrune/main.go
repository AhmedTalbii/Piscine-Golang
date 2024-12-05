package main 

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	if n < 0 {
		return 0
	}
	for index,val := range s {
		if index+1 == n {
			return rune(val)
		}
	}
	return 0
}

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}