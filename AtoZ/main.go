package main

import (
	"github.com/01-edu/z01" 
	"fmt"
)

func main() {
	
	for r := 'a'; r <= 'z'; r++ {
		z01.PrintRune(r) 
	}
	fmt.Printf("\n")
	for r := 'z'; r >= 'a'; r-- {
		z01.PrintRune(r) 
	}
}
