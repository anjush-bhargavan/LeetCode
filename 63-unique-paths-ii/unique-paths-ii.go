func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 || obstacleGrid[0][0] == 1 {
        return 0
    }
    flag := false
    for i := 0 ; i < len(obstacleGrid[0]) ; i++ {
        if obstacleGrid[0][i] != 1 && !flag {
            obstacleGrid[0][i] = 1
        }else{
            obstacleGrid[0][i] = 0
            flag = true
        }
    }
    flag = false
    for i := 1 ; i < len(obstacleGrid) ; i++ {
        if obstacleGrid[i][0] != 1 && !flag {
            obstacleGrid[i][0] = 1
        }else{
            obstacleGrid[i][0] = 0
            flag = true
        }
    }
    i := 1
    for i < len(obstacleGrid) {
        j := 1 
        for j < len(obstacleGrid[i]) {
            if obstacleGrid[i][j] != 1 {
                obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
            }else{
                obstacleGrid[i][j] = 0
            }
            j++
        }
        i++
    }
    return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}