package main

import "fmt"

func Capitalize(s string) string {
	arr := make([]byte, len(s))
	if len(s) > 0 {
		if s[0] >= 'a' && s[0] <= 'z' {
			arr[0] = s[0] - 32
		} else {
			arr[0] = s[0]
		}
	}
	for i := 1; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			if  !(s[i-1] >= 'A' && s[i-1] <= 'Z') &&
				!(s[i-1] >= 'a' && s[i-1] <= 'z') && 
				!(s[i-1] >= '0' && s[i-1] <= '9') { 
				arr[i] = s[i] - 32
			} else {
				arr[i] = s[i]
			}
		} else {
			arr[i] = s[i]
		}
	}
	return string(arr)
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}