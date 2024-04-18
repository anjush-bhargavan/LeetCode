func islandPerimeter(grid [][]int) int {
    count,interCount := 0,0

    for i := 0 ; i < len(grid) ; i++ {
        for j := 0 ; j < len(grid[i]) ; j++ {
            if grid[i][j] == 1 {
                count++
                if i > 0 && grid[i-1][j] == 1 {
                    interCount++
                }
                if i < len(grid)-1 && grid[i+1][j] == 1 {
                    interCount++
                }
                if j > 0 && grid[i][j-1] == 1 {
                    interCount++
                }
                if j < len(grid[i])-1 && grid[i][j+1] == 1 {
                    interCount++
                }
            }
        }
    }
    return (count * 4) - (interCount)
}