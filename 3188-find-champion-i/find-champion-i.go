func findChampion(grid [][]int) int {
    maxSum,champ := 0,0
    for i := 0 ; i < len(grid) ; i++ {
        sum := 0 
        for j := 0 ; j < len(grid[i]) ; j++ {
            if grid[i][j] == 1 {
                sum++
            }
        }
        if maxSum < sum {
            maxSum = sum
            champ = i
        }
    }
    return champ
}