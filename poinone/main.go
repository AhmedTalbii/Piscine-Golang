package main

import "fmt"

func PointOne(i *int) {
	*i++
}


func main() {
	i:=0
	fmt.Printf("var : %v \nadrees : %v \n",i ,&i)
	PointOne(&i)
	fmt.Printf("var : %v \nadrees : %v",i ,&i)


}