package main 

import "fmt"

func quadA(a int,b int) {
	for i := 1 ; i <= b ; i++ {
		for j := 1 ; j <= a ; j++ {
			if (j == 1 && i == 1) || 
			   (j == a && i == 1) || 
			   (j == 1 && i == b) || 
			   (j == a && i == b) {
				fmt.Printf("o")
			} else if j == 1 || j == a {
				fmt.Printf("|")
			} else if i == 1 || i == b {
				fmt.Printf("-")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	quadA(5,3)
}

// o---o
// |   |
// o---o