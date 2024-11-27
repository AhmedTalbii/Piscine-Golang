package main 

import "fmt"

func fibonacciDrawCh(index int) {
	f := 0 
	s := 1
	var res int
	arr := make([]int, index) // This creates a slice with length 5

	// fmt.Println(index)
	
	for i := 0 ; i < index ; i++ {
		res = s + f
		f = s
		s= res
		fmt.Println(res)
		arr[i]= res
		
	}
	fmt.Println(len(arr))
	for j := len(arr)-1 ; j >=0 ; j-- {
		fmt.Println(arr[j])
	}
	for a:=1 ; a <= arr[len(arr)-1] ; a++ {
		for b:=1 ; b <= arr[len(arr)-1] ; b++ {
			fmt.Printf(" 6 ")
		}
		if a <= arr[len(arr)-2] {
			for c:=1 ; c <= arr[len(arr)-2] ; c++ {
				fmt.Printf(" 5 ",)
			}	
		} else if a > arr[len(arr)-2] && a <= arr[len(arr)-2] + arr[len(arr)-5] {
			for e:=1 ; e <= arr[len(arr)-5] ; e++ {
				fmt.Printf(" 2 ")
			}
			for f:=1 ; f <= arr[len(arr)-6] ; f++ {
				fmt.Printf(" %v ", f)
			}
			for g:=1 ; g <= arr[len(arr)-3] ; g++ {
				fmt.Printf(" 4 ")
			}
		} else if a > arr[len(arr)-2] + arr[len(arr)-5] {
			for h:=1 ; h <= arr[len(arr)-4] ; h++ {
				fmt.Printf(" 3 ")
			}
			for n:=1 ; n <= arr[len(arr)-3] ; n++ {
				fmt.Printf(" 4 ")
			}
		}
		fmt.Printf("\n")
	}
	
	
	
}

func main() {
	fibonacciDrawCh(6)
}

// 1  2  3  4  5  6  7  8  9  10  11  12  13  1  2  3  4  5  6  7  8
// 1  2  3  4  5  6  7  8  9  10  11  12  13  1  2  3  4  5  6  7  8
// 1  2  3  4  5  6  7  8  9  10  11  12  13  1  2  3  4  5  6  7  8
// 1  2  3  4  5  6  7  8  9  10  11  12  13  1  2  3  4  5  6  7  8
// 1  2  3  4  5  6  7  8  9  10  11  12  13  1  2  3  4  5  6  7  8
// 1  2  3  4  5  6  7  8  9  10  11  12  13  1  2  3  4  5  6  7  8
// 1  2  3  4  5  6  7  8  9  10  11  12  13  1  2  3  4  5  6  7  8
// 1  2  3  4  5  6  7  8  9  10  11  12  13  1  2  3  4  5  6  7  8 
// 1  2  3  4  5  6  7  8  9  10  11  12  13  2  2  1  5  5  5  5  5 
// 1  2  3  4  5  6  7  8  9  10  11  12  13  2  2  1  5  5  5  5  5  
// 1  2  3  4  5  6  7  8  9  10  11  12  13  3  3  3  5  5  5  5  5  
// 1  2  3  4  5  6  7  8  9  10  11  12  13  3  3  3  5  5  5  5  5  
// 1  2  3  4  5  6  7  8  9  10  11  12  13  3  3  3  5  5  5  5  5  