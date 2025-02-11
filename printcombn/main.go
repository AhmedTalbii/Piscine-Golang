package main

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	arr := make([]int,n)
	for i := 0 ; i < n ; i++ {
		arr[i] = i
	}

	for {
		// print for 
		for i := 0 ; i < len(arr) ; i++ {
			z01.PrintRune(rune(arr[i]+'0'))
		}

		
		// check 
		if arr[0] == 10 - n {
			break
		} else {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}

		// add
		fin := 9
		for i := n-1 ; i >= 0 ; i-- {
			if arr[i] != 9 {
				arr[i]++
			} else {
				fin--
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}