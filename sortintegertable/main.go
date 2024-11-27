package main 
import "fmt"


func SortIntegerTable(table []int) {
	i:= 1
	for i < len(table) {
		if table[i-1] > table[i] {
			temp := table[i-1]
			table[i-1] = table[i]
			table[i] = temp
			i = 1
		}else {
			i++
		}
	}
}

func main() {
	s := []int{5,4,6,2,1,0}
	SortIntegerTable(s)
	fmt.Println(s)
}