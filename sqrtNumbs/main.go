package main

import "fmt"

func sqrtNumbs(nb int) {
	var i int 
	for i=0;i<=nb;i++{
		fmt.Println(i*i)
	}
}

func main() {
	sqrtNumbs(10)
}

//  0   1   4    9    16    25    36  ...    