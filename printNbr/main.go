package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else if n < 0 {
		z01.PrintRune('-')
		n = -n
		rev := 0
		for n > 0 {
			rev = (rev*10)+n%10
			n = n/10
		}
		for rev > 0 {
			z01.PrintRune(rune(rev%10)+'0')
			rev = rev/10
		}
	} else {
		rev := 0
		for n > 0 {
			rev = (rev*10)+n%10
			n = n/10
		}
		for rev > 0 {
			z01.PrintRune(rune(rev%10)+'0')
			rev = rev/10
		}
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}