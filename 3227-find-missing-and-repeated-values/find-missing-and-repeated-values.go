func findMissingAndRepeatedValues(grid [][]int) []int {
    result := []int{}
    valMap := make(map[int]bool)
    for i := 0 ; i < len(grid) ; i++ {
        for j := 0 ; j < len(grid[i]) ; j++ {
            if !valMap[grid[i][j]] {
                valMap[grid[i][j]] = true
            }else{
                result = append(result,grid[i][j])
            }
        }
    }
    for i := 1 ; i <= len(grid)*len(grid) ; i++ {
        if !valMap[i] {
            result = append(result,i)
        }
    } 
    return result
}