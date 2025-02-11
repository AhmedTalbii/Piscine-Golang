package main 

import "fmt"


func Index(s string, toFind string) int {
	if len(s) == 0 || len(s) < len(toFind) {
		return -1
	}
	for i := 0 ; i <= len(s) - len(toFind); i++ {
		if s[i] == toFind[0] {
			res := true
			
			for j := 0 ; j < len(toFind) ; j++ {
				if s[i+j] != toFind[j] {
					res = false 
				}
			}
			if res {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!llca", "allc"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}