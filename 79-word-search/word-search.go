func exist(board [][]byte, word string) bool {
    row, col := len(board), len(board[0])
    indexMap := make(map[[2]int]bool)

    var dfs func(r, c, i int) bool
    dfs = func(r, c, i int) bool {
        if i == len(word) {
            return true
        }
        if r < 0 || c < 0 || r >= row || c >= col || word[i] != board[r][c] || indexMap[[2]int{r, c}] {
            return false
        }
        indexMap[[2]int{r, c}] = true
        defer func() { indexMap[[2]int{r, c}] = false }()
        return dfs(r+1, c, i+1) || dfs(r-1, c, i+1) || dfs(r, c+1, i+1) || dfs(r, c-1, i+1)
    }

    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }
    return false
}

