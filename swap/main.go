package main 

import "fmt"

func Swap(a *int, b *int) {
	o:=*a
	*a=*b
	*b=o
}

func main() {
	a:=0
	b:=1
	Swap(&a,&b)
	fmt.Printf("%v\n%v",a,b)
}