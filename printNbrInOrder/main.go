package main 

import "fmt"

func PrintNbrInOrder(n int) {
	if n == 0 {
		fmt.Printf("%v",0)
		return
	}
	arr := make([]int,0)
	for n > 0 {
		arr = append(arr,n%10)
		n = n/10
	}
	for i := 1 ; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			temp := arr[i]
			arr[i] = arr[i-1]
			arr[i-1]= temp
			i = 1
		}
	}
	for i := 0 ; i < len(arr) ; i++ {
		fmt.Printf("%v",arr[i])
	}
}

func main() {
	// PrintNbrInOrder(321)
	// PrintNbrInOrder(0)
	// PrintNbrInOrder(321)
	// PrintNbrInOrder(9876543210)
	// PrintNbrInOrder(21110)
	// PrintNbrInOrder(987654322222222222)
	PrintNbrInOrder(0)
	// PrintNbrInOrder(3210)
	fmt.Println()
}