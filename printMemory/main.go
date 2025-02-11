package main 

import (
	"github.com/01-edu/z01"
)

func base(arr [10]byte) {
	base := "0123456789abcdef"
	counter := 0
	for _,val := range arr {
		f :=  val/16
		g :=  val%16
		z01.PrintRune(rune(f)+'0')
		z01.PrintRune(rune(base[g]))
		z01.PrintRune(' ')
		counter++
		if counter == 4 {
			z01.PrintRune('\n')
			counter = 0
		}
	}
	z01.PrintRune('\n')
}

func PrintMemory(arr [10]byte) {
	base(arr)
	for _,char := range arr {
		z01.PrintRune(rune(char))
		if char < 32 || char > 126 {
			z01.PrintRune('.')
		}
	}
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}