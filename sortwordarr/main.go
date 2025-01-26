package main 

import "fmt"


func SortWordArr(a []string) {
	for i := 0 ; i < len(a)-1 ; i++ {
		for j := i+1 ; j < len(a) ; j++ {
			if a[i] > a[j] {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}
}

func main() {
	arr := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(arr)
	fmt.Println(arr)
}