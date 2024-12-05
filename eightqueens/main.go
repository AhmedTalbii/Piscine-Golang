package main

import "fmt"

// n is the size of the board and number of queens
const n = 8

var matrix [n][n]string
var solutions []string

// Check if it's safe to place a queen at matrix[row][col]
func isSafe(row, col int) bool {
	// Check this column on upper rows
	for i := 0; i < row; i++ {
		if matrix[i][col] == "Q" {
			return false
		}
	}

	// Check upper left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if matrix[i][j] == "Q" {
			return false
		}
	}

	// Check upper right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if matrix[i][j] == "Q" {
			return false
		}
	}

	return true
}

// Solve N-Queens using backtracking
func solve(row int) {
	if row == n {
		// All queens placed successfully, store the solution
		storeSolution()
		return
	}

	for col := 0; col < n; col++ {
		if isSafe(row, col) {
			// Place a queen
			matrix[row][col] = "Q"

			// Recursively place the rest of the queens
			solve(row + 1)

			// Backtrack: Remove the queen
			matrix[row][col] = "."
		}
	}
}

// Store the current matrix configuration as a solution
func storeSolution() {
	var solution string
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			solution += matrix[i][j]
		}
		solution += "\n"
	}
	solutions = append(solutions, solution)
}

func main() {
	// Initialize the board
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = "."
		}
	}

	// Solve the N-Queens problem
	solve(0)

	// Print all solutions
	fmt.Printf("Total solutions: %d\n", len(solutions))
	for i, solution := range solutions {
		fmt.Printf("Solution %d:\n%v\n", i+1, solution)
	}
}
