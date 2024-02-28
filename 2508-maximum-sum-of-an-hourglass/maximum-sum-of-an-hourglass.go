func maxSum(grid [][]int) int {
    maxSum := 0

    for i := 0 ; i < len(grid) ; i++ {
        if i + 3 <= len(grid) {
            for j := 0 ; j < len(grid[i]) ; j++ {
                if j + 3 <= len(grid[i]) {
                    sum := grid[i][j]+grid[i][j+1]+grid[i][j+2]+grid[i+1][j+1]+ grid[i+2][j]+grid[i+2][j+1]+grid[i+2][j+2]
                    if sum > maxSum {
                        maxSum = sum
                    }
                }else{
                    break
                } 
            }
        }else{
            break
        }
    }
    return maxSum
}