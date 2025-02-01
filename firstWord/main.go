package main

import "fmt"

func FirstWord(s string) string {
	res := ""
	started := false 

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' {
			res = res + string(s[i])
			started = true
		} else if started {
			break
		}
	}

	return res + "\n"
}

func main() {
	fmt.Print(FirstWord("hello world"))             // Expected output: "hello\n"
	fmt.Print(FirstWord("   leading spaces"))       // Expected output: "leading\n"
	fmt.Print(FirstWord("trailing spaces   "))      // Expected output: "trailing\n"
	fmt.Print(FirstWord("   multiple   spaces   ")) // Expected output: "multiple\n"
	fmt.Print(FirstWord("singleword"))              // Expected output: "singleword\n"
	fmt.Print(FirstWord(""))                        // Expected output: "\n" (empty string case)
	fmt.Print(FirstWord("   "))                     // Expected output: "\n" (only spaces)
	fmt.Print(FirstWord("a"))                       // Expected output: "a\n"
	fmt.Print(FirstWord("  a  "))                   // Expected output: "a\n"
	fmt.Print(FirstWord("word1 word2 word3"))       // Expected output: "word1\n"
	fmt.Print(FirstWord(" 123 456 789"))            // Expected output: "123\n"
	fmt.Print(FirstWord("\t tabbed text"))          // Expected output: "tabbed\n"
	fmt.Print(FirstWord("punctuation? yes!"))       // Expected output: "punctuation?\n"
	fmt.Print(FirstWord("\nnewline at start"))      // Expected output: "newline\n"
	fmt.Print(FirstWord("mix\t of spaces and tabs"))// Expected output: "mix\n"
}