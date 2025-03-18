package main

import "fmt"


func intToBinary(num int) []int {
    rev := make([]int, 0)
    for num > 0 {
        rev = append(rev, num%2)
        num = num / 2
    }
    
    res := make([]int, 0)
    for i := len(rev)-1 ; i >= 0; i-- {
        res = append(res, rev[i])
    }
    
    if len(res) == 0 {
        return []int{0}
    }
    
    return res
}

func power(a, b int) int {
    res := 1
    for i := 0 ; i < b ; i++ {
        res = res * a
    }
    return res
}

func binaryToInt(arr []int) int {
    res := 0
    counter := 0
    for i := len(arr)-1 ; i >= 0 ; i-- {
        res = res + arr[i] * power(2,counter)
        counter++
    }
    return res
}


func compare(arr1 []int, arr2 []int) []int {
    maxLen := len(arr1)
    if len(arr2) > maxLen {
        maxLen = len(arr2)
    }

    for len(arr1) < maxLen {
        arr1 = append([]int{0}, arr1...)
    }
    for len(arr2) < maxLen {
        arr2 = append([]int{0}, arr2...)
    }

    res := make([]int, maxLen)
    for i := 0; i < maxLen; i++ {
        if arr1[i] == arr2[i] {
            res[i] = 0
        } else {
            res[i] = 1
        }
    }
    return res
}

func xor(a, b int) int {
    arr1 := intToBinary(a)
    arr2 := intToBinary(b)
    res := compare(arr1, arr2)
    return binaryToInt(res)
}
	
func Unmatch(a []int) int {
	result := 0
	for _, num := range a {
		result = xor(result,num)
	}
    if result == 0 {
        return -1
    }
	return result
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
