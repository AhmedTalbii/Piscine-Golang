package main 

import "fmt"

func Max(arr []int) int {
	if  len(arr) == 0 {
		return 0
	}
	var maxVal,a,i int
	var b int = arr[0]
	for i=1 ; i < len(arr) ; i++ {
	   a = arr[i] 
	   if a > b {
		   b = a
	   }
	}
	maxVal = b
	return maxVal   
}

func main() {
		// Test Case 1: Normal Case with positive integers
		a1 := []int{23, 123, 1, 11, 55, 93}
		max1 := Max(a1)
		fmt.Println(max1) // Expected output: 123
	
		// Test Case 2: Slice with a single element
		a2 := []int{42}
		max2 := Max(a2)
		fmt.Println(max2) // Expected output: 42
	
		// Test Case 3: Empty slice
		a3 := []int{}
		max3 := Max(a3)
		fmt.Println(max3) // Expected output: 0
	
		// Test Case 4: All negative integers
		a4 := []int{-1, -23, -5, -42}
		max4 := Max(a4)
		fmt.Println(max4) // Expected output: -1
	
		// Test Case 5: Slice with only one negative number
		a5 := []int{-42}
		max5 := Max(a5)
		fmt.Println(max5) // Expected output: -42
	
		// Test Case 6: Slice with repeated values
		a6 := []int{5, 5, 5, 5}
		max6 := Max(a6)
		fmt.Println(max6) // Expected output: 5

		// Test Case 7: Slice with mixed positive and negative numbers
		a7 := []int{-1, 2, 0, 100, -5}
		max7 := Max(a7)
		fmt.Println(max7) // Expected output: 100
		fmt.Println(Max([]int{2147483647, -2147483648, 0, 99999999})) 
		fmt.Println(Max([]int{-10, -10, -10, -10})) 
		fmt.Println(Max([]int{-1000, 500, -1000, 500, -1000, 500, 999})) 
		fmt.Println(Max([]int{-5, -1, -100, 0})) 
}