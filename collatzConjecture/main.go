package main 

import "fmt"

func CollatzCountdown(start int) int {
	steps := 0
	if start <= 0 {
		return -1
	}
	for start != 1 {
		if start%2 == 0 {
			start = start/2
		} else {
			start = (start*3) + 1
		}
		steps++
	}
	return steps
}

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}