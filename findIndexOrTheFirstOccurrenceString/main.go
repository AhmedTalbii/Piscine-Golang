package main 

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true
		for j := 0; j < len(needle); j++ {
			fmt.Println("i: ",i," j: ",j," haystack[i+j] ",string(haystack[i+j])," needle[j]: ",string(needle[j]))
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}


func main() {
	// fmt.Println(strStr("sadbutsad","sad"))
	fmt.Println(strStr("mississippi","issip"))
	// fmt.Println(strStr("leetcode","leeto"))
	// fmt.Println(strStr("aaaaa","bba"))
	// fmt.Println(strStr("mississippi","a"))

}