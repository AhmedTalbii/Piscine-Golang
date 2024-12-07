package main 

import "fmt"

func ToLower(s string) string {
	arr := make([]byte,len(s))
	for index,char := range s {
		if char >= 'A' && char <= 'Z' {
			arr[index] = byte(char) + ('a'-'A')
		} else {
			arr[index] = byte(char)
		}
	}
	return string(arr)
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}