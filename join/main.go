package main 
import "fmt"

func Join(strs []string, sep string) string {
	res := strs[0]
	for i := 1 ; i < len(strs) ; i++ {
		res = res + sep + strs[i]
	}
	return res
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}