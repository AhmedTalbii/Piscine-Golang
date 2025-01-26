package main 

import "fmt"

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _,v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func CountIf(f func(string) bool, tab []string) int {
	res := 0 
	for _,v := range tab {
		if f(v) {
			res++
		}
	}
	return res
}

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}