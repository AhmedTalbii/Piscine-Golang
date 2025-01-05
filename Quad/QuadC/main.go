package main 

import "fmt"

func quadA(a int,b int) {
	for i := 1 ; i <= b ; i++ {
		for j := 1 ; j <= a ; j++ {
			if j == 1 && i == 1 {
				fmt.Printf("A")
			} else if j == a && i == 1 {
				fmt.Printf("A")
			} else if j == 1 && i == b {
				fmt.Printf("C")
 		    } else if j == a && i == b {
				fmt.Printf("C")
			} else if j == 1 || j == a {
				fmt.Printf("B")
			} else if i == 1 || i == b {
				fmt.Printf("B")
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

// ABBBA
// B   B
// CBBBC