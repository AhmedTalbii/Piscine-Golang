package main
import "fmt"

func StrRev(s string) string{
	l:= len(s)
	new := make([]byte, l) 
	j:=0
	for i:= len(s)-1 ; i>=0; i-- {
		
		new[j] = s[i]
		j++
		
	}
	return string(new)
}

func main() {
	s := "Hello World!"
	s = StrRev(s) 
	fmt.Println(s)
}