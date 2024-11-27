package main

import "fmt"
func UltimatePointOne(n ***int) {
	***n++
}
func main() {
	n:=0
	a:=&n
	b:=&a
	

	UltimatePointOne(&b)
	fmt.Printf("%v",n)

}