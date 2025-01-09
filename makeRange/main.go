package main

import "fmt"

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arr := make([]int,max-min)
	counter := 0
	for i := min ; i < max ; i++ {
		arr[counter] = i
		counter++
	}
	return arr
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}