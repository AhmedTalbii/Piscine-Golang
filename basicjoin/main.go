package main 

import "fmt"

func BasicJoin(elems []string) string {
	res := ""
	for i:=0;i<len(elems);i++{
		res = res + elems[i]
	}
	return res
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}