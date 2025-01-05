package main 

import "fmt"

func SortIntegerTable(s []int) {
	i := 1
	for i < len(s) {
		if s[i-1] > s[i] {
			temp := s[i-1]
			s[i-1] = s[i]
			s[i] = temp
			i = 1
		} else {
			i++
		}
	}
}

func main() {
	s := []int{5,4,6,2,1,0}
	SortIntegerTable(s)
	fmt.Println(s)
}