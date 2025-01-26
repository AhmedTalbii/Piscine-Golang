package main 

import (
	"fmt"
	"os"
)

func calcul(a string, b string, c string) int {

}

func applyFunc(f func(a int, b , c int) int , arr []string) int {
	if len(arr) < 4 && len(arr) > 0{
		res := calcul()
	}
}

func main() {
	arr := os.Args[1:]
	res := applyFunc(calcul, arr)
	
}