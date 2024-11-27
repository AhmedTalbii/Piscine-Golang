package main 

import "fmt"

func DivMod(a int, b int, div *int, mod *int) {
	*div = a/b
	*mod = a%b
}

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Printf("%v\n",div)
	fmt.Printf("%v",mod)
}