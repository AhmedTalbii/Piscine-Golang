package main
import "fmt"

func Sqrt(nb int) int {
	for i:=1 ; i<=nb; i++ {
		if i * i == nb {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Printf("%v\n",Sqrt(4))
	fmt.Printf("%v",Sqrt(3))
}