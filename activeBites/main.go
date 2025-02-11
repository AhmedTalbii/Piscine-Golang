package main 

import "fmt"

func ActiveBits(n int) int {
	res := 0
	for n != 0 {
		if n%2 == 1 {
			res++
		}
		n = n / 2
	}
	return res
}

func main() {
	fmt.Println(ActiveBits(7))
	fmt.Println(ActiveBits(2147483647)) // Expected Output: 31 (Binary representation of 2147483647)
	fmt.Println(ActiveBits(8)) // Expected Output: 1 (Binary of 8 is 1000)
	fmt.Println(ActiveBits(16)) // Expected Output: 1 (Binary of 16 is 10000)
	
}