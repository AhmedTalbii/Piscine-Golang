package main 

import "fmt"


func Concat(str1 string, str2 string) string {
	concat := ""
	concat = str1 + str2
	return concat
}

// func Concat(str1 string, str2 string) string {
// 	arr := make([]byte, len(str1)+len(str2))
// 	for i:=0 ; i < len(str1); i++{
// 		arr[i] = str1[i]
// 	}
// 	counter := 0
// 	for j:= len(str1) ; j < len(str1)+len(str2) ;j++{
// 		arr[j] = str2[counter]
// 		counter++
// 	}
// 	return string(arr)
// }

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))
}