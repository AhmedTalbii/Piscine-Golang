package main

import (
	"fmt"
	"os"
)

func checkRow(matrix [9][9]string, i int, arr []string) {
	for j := 0; j < 9; j++ {
		for k := 0; k < len(arr); k++ {
			if arr[k] == matrix[i][j] {
				arr[k] = "E"
			}
		}
	}
}

func checkCol(matrix [9][9]string, i int, arr []string) {
	for j := 0; j < 9; j++ {
		for k := 0; k < len(arr); k++ {
			if arr[k] == matrix[j][i] {
				arr[k] = "E"
			}
		}
	}
}

func checkSq(matrix [9][9]string, i int, j int, arr []string) {
	// Top-left square
	if i >= 0 && i <= 2 && j >= 0 && j <= 2 {
		for row := 0; row <= 2; row++ {
			for col := 0; col <= 2; col++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] == matrix[row][col] {
						arr[k] = "E"
					}
				}
			}
		}
	}
	// Top-middle square
	if i >= 0 && i <= 2 && j >= 3 && j <= 5 {
		for row := 0; row <= 2; row++ {
			for col := 3; col <= 5; col++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] == matrix[row][col] {
						arr[k] = "E"
					}
				}
			}
		}
	}
	// Top-right square
	if i >= 0 && i <= 2 && j >= 6 && j <= 8 {
		for row := 0; row <= 2; row++ {
			for col := 6; col <= 8; col++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] == matrix[row][col] {
						arr[k] = "E"
					}
				}
			}
		}
	}
	// Middle-left square
	if i >= 3 && i <= 5 && j >= 0 && j <= 2 {
		for row := 3; row <= 5; row++ {
			for col := 0; col <= 2; col++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] == matrix[row][col] {
						arr[k] = "E"
					}
				}
			}
		}
	}
	// Middle-middle square
	if i >= 3 && i <= 5 && j >= 3 && j <= 5 {
		for row := 3; row <= 5; row++ {
			for col := 3; col <= 5; col++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] == matrix[row][col] {
						arr[k] = "E"
					}
				}
			}
		}
	}
	// Middle-right square
	if i >= 3 && i <= 5 && j >= 6 && j <= 8 {
		for row := 3; row <= 5; row++ {
			for col := 6; col <= 8; col++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] == matrix[row][col] {
						arr[k] = "E"
					}
				}
			}
		}
	}
	// Bottom-left square
	if i >= 6 && i <= 8 && j >= 0 && j <= 2 {
		for row := 6; row <= 8; row++ {
			for col := 0; col <= 2; col++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] == matrix[row][col] {
						arr[k] = "E"
					}
				}
			}
		}
	}
	// Bottom-middle square
	if i >= 6 && i <= 8 && j >= 3 && j <= 5 {
		for row := 6; row <= 8; row++ {
			for col := 3; col <= 5; col++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] == matrix[row][col] {
						arr[k] = "E"
					}
				}
			}
		}
	}
	// Bottom-right square
	if i >= 6 && i <= 8 && j >= 6 && j <= 8 {
		for row := 6; row <= 8; row++ {
			for col := 6; col <= 8; col++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] == matrix[row][col] {
						arr[k] = "E"
					}
				}
			}
		}
	}
}

func main() {
	arr := os.Args[1:]
	if len(arr) != 9 {
		fmt.Println("Error")
		return
	}
	res := [9][9]string{}
	stop := false
	for i, char := range arr {
		if len(char) != 9 {
			stop = true
			break
		} else {
			for j, v := range char {
				res[i][j] = string(v)
			}
		}
	}
	if stop {
		fmt.Println("Error")
		return
	}

	for {
		filled := 0
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if res[i][j] == "." {
					poss := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
					checkRow(res, i, poss)
					checkCol(res, i, poss)
					checkSq(res, i, j, poss)
					count := 0
					val := ""
					for _, num := range poss {
						if num != "E" {
							count++
							val = num
						}
					}
					if count == 1 {
						res[i][j] = val
						filled++
					}
				}
			}
		}
		if filled == 0 {
			break
		}
	}

	// Check if solved
	complete := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if res[i][j] == "." {
				complete = false
				break
			}
		}
	}

	// Print result
	if complete {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Printf("%s", res[i][j])
			}
			fmt.Println()
		}
	}
}