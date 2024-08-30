# Backtracking Algorithms Documentation

## Overview: What is Backtracking?
**Backtracking** is a problem-solving technique that uses recursion to explore all possible solutions for a given problem. It systematically builds candidates for the solution and abandons a candidate ("backtracks") as soon as it determines that the candidate cannot possibly lead to a valid solution.

---

### Q: N-Queen Problem
**Problem**: Place N queens on an N×N chessboard such that no two queens threaten each other. This means that no two queens can be in the same row, column, or diagonal.

**Solution**: 

- Use a backtracking approach to place queens one by one in different columns, starting from the leftmost column. 
- For each placement, check if the current position is safe using an `isSafe` function. 
- If placing the queen leads to a solution, return true. If not, backtrack and try the next possible position.

```go
func nQueen(board [][]int, row int) bool {
    if row == len(board) { // Base case: All queens are placed
        return true
    }

    cols := len(board[0])
    for col := 0; col < cols; col++ { // Try placing queen in all columns
        if isSafe(board, row, col) { // Check if it's safe to place the queen
            board[row][col] = 1 // Place the queen
            if nQueen(board, row+1) { // Recur to place the rest of the queens
                return true
            }
            board[row][col] = 0 // Backtrack: Remove the queen
        }
    }
    return false // No safe position found, backtrack to the previous row
}
```

---

### Q: Sudoku Solver
**Problem**: Solve a given 9x9 Sudoku puzzle by filling the empty cells. A valid Sudoku solution must satisfy all the constraints of the game: each number 1-9 must appear exactly once in each row, column, and 3x3 subgrid.

**Solution**: 

- Use backtracking to fill the empty cells one by one. 
- For each empty cell, try placing each number from 1 to 9. 
- Check if the placement is valid using a function to ensure the number doesn't violate Sudoku rules.
- If a valid placement is found, move to the next empty cell.
- If no valid number can be placed, backtrack to the previous cell and try a different number.

```go
package main

import "fmt"

// Function to solve Sudoku using backtracking
func solveSudoku(board [][]int) bool {
    row, col, empty := findEmpty(board)
    if !empty {
        return true // No empty cell left, puzzle is solved
    }

    for num := 1; num <= 9; num++ {
        if isValid(board, row, col, num) { // Check if num is valid
            board[row][col] = num
            if solveSudoku(board) { // Recur to solve the next empty cell
                return true
            }
            board[row][col] = 0 // Backtrack if num doesn't lead to a solution
        }
    }
    return false
}

// Helper function to find the next empty cell
func findEmpty(board [][]int) (int, int, bool) {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == 0 {
                return i, j, true
            }
        }
    }
    return 0, 0, false
}

// Helper function to check if placing num in board[row][col] is valid
func isValid(board [][]int, row, col, num int) bool {
    for i := 0; i < 9; i++ {
        if board[row][i] == num || board[i][col] == num || board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
            return false
        }
    }
    return true
}

func main() {
    board := [][]int{
        {5, 3, 0, 0, 7, 0, 0, 0, 0},
        {6, 0, 0, 1, 9, 5, 0, 0, 0},
        {0, 9, 8, 0, 0, 0, 0, 6, 0},
        {8, 0, 0, 0, 6, 0, 0, 0, 3},
        {4, 0, 0, 8, 0, 3, 0, 0, 1},
        {7, 0, 0, 0, 2, 0, 0, 0, 6},
        {0, 6, 0, 0, 0, 0, 2, 8, 0},
        {0, 0, 0, 4, 1, 9, 0, 0, 5},
        {0, 0, 0, 0, 8, 0, 0, 7, 9},
    }

    if solveSudoku(board) {
        fmt.Println("Solved Sudoku:")
        printBoard(board)
    } else {
        fmt.Println("No solution exists.")
    }
}

// Function to print the Sudoku board
func printBoard(board [][]int) {
    for _, row := range board {
        fmt.Println(row)
    }
}
```

---

This documentation provides an overview of backtracking and its application in solving complex problems like the N-Queen problem and Sudoku. The code examples illustrate how backtracking can be implemented in Go, with detailed explanations to help you understand the underlying concepts.