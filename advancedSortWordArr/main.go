package main
 
import "fmt"

func Compare(a, b string) int {
	if a > b {
		return 1
	}
	if b > a {
		return -1
	}
	return 0
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0 ; i < len(a)-1 ; i++ {
		for j := i+1 ; j < len(a) ; j++ {
			num := f(a[i],a[j])
			if num > 0 {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		} 
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)
	fmt.Println(result)
}