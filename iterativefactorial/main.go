package main 

import "fmt"

func IterativeFactorial(nb int) int {
    if nb <= 0 || nb >= 20 {
        return 0
    }
    result := 1
    for i := 1; i <= nb; i++ {
        result = result * i
    }
    return result
}


func main() {
	age := 4
	it := IterativeFactorial(age)
	fmt.Println(it)
}