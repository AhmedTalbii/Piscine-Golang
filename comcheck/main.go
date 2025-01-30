package main 

import (
	"os"
	"fmt"
)

func main() {
	arr := os.Args[1:]
	for _,str := range arr {
		if str == "01" || str == "galaxy" || str == "galaxy 01" {
			fmt.Printf("Alert!!!\n")
			return
		}
	}
}