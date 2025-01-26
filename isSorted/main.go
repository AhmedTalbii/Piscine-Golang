package main 

import "fmt"

func f(a, b int) int {
	// min to max
	if a+1 == b {
		return 1
	}
	// max to min 
	if a == b+1 {
		return 2
	}
	// equal 
	if a == b {
		return 3
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

 	mitMa := 0
	matMi := 0
	equal := 0

	for i := 1 ; i < len(a) ;i++ {
		// min to max
		if f(a[i-1],a[i]) == 1 {
			mitMa++
		}
		// max to min 
		if f(a[i-1],a[i]) == 2 {
			matMi++
		}
		// equal
		if f(a[i-1],a[i]) == 3 {
			equal++
		}
	}
	if mitMa+1 == len(a) {
		return true 
	}
	if matMi+1 == len(a) {
		return true
	}
	if equal+1 == len(a) {
		return true
	}
	return false
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{6, 5, 4, 3, 2, 1}
	a4 := []int{7, 7, 7, 7}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)
	result3 := IsSorted(f, a3)
	result4 := IsSorted(f, a4)


	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}