func uniquePaths(m int, n int) int {
    row := make([]int,n)
    for i := range n {
        row[i] = 1
    }

    for _ = range m - 1 {
        newRow := make([]int,n)
        newRow[n-1] = 1
        for j := n - 2 ; j >= 0 ; j-- {
            newRow[j] = newRow[j+1] + row[j]
        }
        row = newRow
    }
    return row[0]
}