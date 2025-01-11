package main
import "fmt"

func Split(s, sep string) []string {
	res := make([]string,0)
	word := ""
	for i := 0 ; i < len(s)-len(sep) ; i++ {
		// fmt.Println(string(s[i:i+len(sep)]))
		// fmt.Println(string(sep[0:len(sep)]))
		if s[i:i+len(sep)] == sep[0:len(sep)] {
			res = append(res,word)
			word = ""
		} else {
			word = word + string(s[i])
		}
	}
	if s[len(s)-len(sep):len(s)] == sep[0:len(sep)] {
		res = append(res,word)
		res = append(res,"")
	} else {
		word = word + string(s[len(s)-len(sep):len(s)])
		res = append(res,word)
	}
	return res
}

// func Split(s, sep string) []string {
// 	resB := make([]string, 0)
// 	word := ""
// 	counter := len(sep)
// 	for i := 0 ; i < len(s)-len(sep) ; i++ {
// 		for j := 0 ; j < len(sep) ; j++ {
// 			if s[i+j] == sep[j] {
// 				counter--
// 				if counter == 0 {
// 					resB = append(resB, word)
// 					word = ""
// 					counter = len(sep)
// 				}
// 			}
// 		}
// 		word = word + string(s[i])
// 		if i == len(s)-len(sep)-1 {
// 			for j := len(s)-len(sep) ; j < len(s) ; j++ {
// 				word = word + string(s[j])
// 			}
// 			resB = append(resB, word)
// 		}
// 	}
// 	resA := make([]string, 0)
// 	counterA := len(sep)
// 	wordA := ""
// 	match := false
// 	for _,v := range resB {
// 		for i := 0 ; i < len(sep) ; i++ {
// 			if v[i] == sep[i] {
// 				counterA--
// 				if counterA == 0 {
// 					match = true 
// 					counter = len(sep)
// 				}
// 			}
// 		}
// 		if !match {
// 			resA = append(resA, v)
// 		} else {
// 			for i := len(sep) ; i < len(v) ; i++ {
// 				wordA = wordA + string(v[i])
// 			}
// 			resA = append(resA, wordA)
// 			wordA = ""
// 		}
// 	}
// 	return resA
// }

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
	s = "one;two;three;four"
	fmt.Printf("%#v\n", Split(s, ";"))
	s = "no-separator-here"
	fmt.Printf("%#v\n", Split(s, "."))
	s = "xxaxxbxxcxxdxx"
	fmt.Printf("%#v\n", Split(s, "xx"))
	s = "////"
	fmt.Printf("%#v\n", Split(s, "/"))
}