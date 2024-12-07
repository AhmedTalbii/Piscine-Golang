package main
import "fmt"

func ToUpper(s string) string {
	arr := make([]byte, len(s))
	for index,val := range s {
		if val >= 'a' && val <= 'z' {
			val = val - ('a'-'A')
			arr[index] = byte(val)
		} else {
			arr[index] = byte(val)
		}
	}
	return string(arr)
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}