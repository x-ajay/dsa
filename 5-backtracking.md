Find optimal problem solution from recuression

Q: N-Queen Problem

```go
func nQueen(board [][]int,row int) bool{
    if row == len(board) {
        return true
    }
    cols := len(board[0])
    for col := 0 ; i < cols ; col++ {
        if isSafe(board,row,col) {
            board[row][col] = 1
            if nQueen(board,row+1){
                return true
            }
            board[row][col] = 0 // back-tracking 
        }
    }
    return false
}
```

---

Q: Sudoku Solver 

```go
```