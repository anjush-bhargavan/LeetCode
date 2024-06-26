func setZeroes(matrix [][]int)  {
    rowMap := make(map[int]bool)
    colMap := make(map[int]bool)

    for i := 0 ; i < len(matrix) ; i++ {
        for j := 0 ; j < len(matrix[i]) ; j++ {
            if matrix[i][j] == 0 {
                rowMap[i] = true
                colMap[j] = true
            }
        }
    }
    for i := 0 ; i < len(matrix) ; i++ {
        for j := 0 ; j < len(matrix[i]) ; j++ {
            if rowMap[i] {
                matrix[i][j] = 0
            }
            if colMap[j] {
                matrix[i][j] = 0
            }
        }
    }
}