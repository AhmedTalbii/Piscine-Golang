package main 

import "fmt"

func Abort(a, b, c, d, e int) int {
	arr := []int{a,b,c,d,e}
	for i := 1 ; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			temp := arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = temp
			i=1
		}
	}
	return arr[len(arr)/2]
}

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}