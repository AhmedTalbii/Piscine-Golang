package main 

import "fmt"

func Unmatch(a []int) int {
	hashMap := make(map[int]int)
	for _,char := range a {
		hashMap[char]++
	}
	for key, val := range hashMap {
		if val == 1 {
			return key
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}