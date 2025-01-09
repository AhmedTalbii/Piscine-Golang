package main 

import "fmt"

func ConcatParams(args []string) string {
	res := ""
	for i := 0 ; i < len(args) ; i++ {
		res = res + args[i]
		if i < len(args)-1 {
			res = res + "\n"
		}
	}
	return res
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}