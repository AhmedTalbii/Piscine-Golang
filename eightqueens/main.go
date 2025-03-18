package main

import "fmt"

const n = 8

var board = [n][n]int{}
var numSels = 0

func underAtack(x int,y int,z int) bool {
	x,y = x + z , y -1
	for y >= 0 && x >= 0 && x < n {
		if board[y][x] == 1 {
			return true
		}
		x = x + z
		y = y - 1
	}
	return false
}

func solve(y int) {
	if y == n {
		numSels = numSels + 1
		for _ , row := range board {
			fmt.Println(row)
		}
		fmt.Println()
	}
	for x := range [n]int{} {
		if !(underAtack(x ,y ,0)||underAtack(x ,y ,-1)||underAtack(x ,y ,1)) {
			board[y][x] = 1
			solve(y+1)
			board[y][x] = 0
		}
	}
}

func main() {	
	solve(0)
	fmt.Printf("found %d solutions \n", numSels)
}
