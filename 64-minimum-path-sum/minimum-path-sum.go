func minPathSum(grid [][]int) int {
    for i := len(grid[0])-2 ; i >= 0 ; i-- {
        grid[len(grid)-1][i] += grid[len(grid)-1][i+1]
    }
    for i := len(grid)-2 ; i >= 0 ; i-- {
        grid[i][len(grid[0])-1] += grid[i+1][len(grid[0])-1]
    }

    i:= len(grid)-2
    for  i >= 0 {
        j := len(grid[0])-2
        for j >= 0 {
            grid[i][j] += min(grid[i+1][j],grid[i][j+1])
            j--
        }
        i--
    }
    
    return grid[0][0]
}