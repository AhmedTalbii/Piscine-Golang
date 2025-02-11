package main 

import "fmt"

func even(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false 
}

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	res := 0
	for start != 1 {
		if even(start) {
			start = start/2
			res++
		} else {
			start = start * 3 + 1
			res++
		}
	}
	return res
}

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}