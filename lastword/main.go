package main 

import "fmt"

func LastWord(str string) string {
	if len(str) == 0 {
		return ""
	}
	res := ""
	for i := 1 ; i < len(str); i++ {
		if (str[i] != ' ' && str[i-1] == ' ' || str[i-1] == '\t') {
			res = ""
		} else {
			if str[i-1] != ' ' && len(res) == 0 {
				res = res + string(str[i-1])
			}
			res = res + string(str[i])
		}
	}
	res = res + "\n"
	return res
}

func main() {
	fmt.Print(LastWord("hello world"))               // Expected output: "world$"
	fmt.Print(LastWord("   leading spaces"))         // Expected output: "spaces$"
	fmt.Print(LastWord("trailing spaces   "))        // Expected output: "spaces$"
	fmt.Print(LastWord("   multiple   spaces   "))   // Expected output: "spaces$"
	fmt.Print(LastWord("singleword"))                // Expected output: "singleword$"
	fmt.Print(LastWord(""))                          // Expected output: "$" (empty string case)
	fmt.Print(LastWord("   "))                       // Expected output: "$" (only spaces)
	fmt.Print(LastWord("a"))                         // Expected output: "a$"
	fmt.Print(LastWord("  a  "))                     // Expected output: "a$"
	fmt.Print(LastWord("word1 word2 word3"))         // Expected output: "word3$"
	fmt.Print(LastWord("multiple\nlines here"))      // Expected output: "here$"
	fmt.Print(LastWord("punctuation? yes!"))         // Expected output: "yes!$"
	fmt.Print(LastWord("123 456 789"))               // Expected output: "789$"
	fmt.Print(LastWord("tab\tseparated\twords"))     // Expected output: "words$" (if tabs are treated as spaces)	
}