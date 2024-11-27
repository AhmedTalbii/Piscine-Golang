package main

import "fmt"

func max(arr []int)(int, int) {

var maxVal,minVal,a,i int
var ln int = len(arr)
var b int = arr[0]
   for i=1 ; i<ln; i++ {
	a = arr[i]
	if a > b {
		b = a
	}
   }
   maxVal = b
   for i=1 ; i<ln; i++ {
	a = arr[i]
	if a < b {
		b = a
	}
   }
   minVal = b
   return maxVal,minVal
}

func main() {
   var arr = []int{3,1,150,5,80}
   var maxVal, minVal int
	
	maxVal, minVal = max(arr)
   fmt.Println("Max:", maxVal)
   fmt.Println("Min:", minVal)
}