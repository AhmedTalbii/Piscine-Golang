package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	t := *a
	*a = *a / *b
	*b = t % *b
}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Printf("%v \n%v", a, b)
}
