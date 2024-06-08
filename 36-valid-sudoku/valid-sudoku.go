func isValidSudoku(board [][]byte) bool {
  colMap := make([]map[byte]bool,len(board))
  gridMap := make([]map[byte]bool,len(board))
  for i := range 9 {
    colMap[i] = make(map[byte]bool)
    gridMap[i] = make(map[byte]bool)
  }

  for i := 0 ; i < len(board) ; i++ {    
    rowMap := make(map[byte]bool)
    for j := 0 ; j < len(board[i]) ; j++ {
        if board[i][j] != '.' {

            if !rowMap[board[i][j]] {
                rowMap[board[i][j]] = true
            }else{
                fmt.Println("here",i)
                return false
            }

            if !colMap[j][board[i][j]] {
                colMap[j][board[i][j]] = true
            }else{
                return false
            }

            if !gridMap[(i/3)*3+(j/3)][board[i][j]] {
                gridMap[(i/3)*3+(j/3)][board[i][j]] = true
            }else {
                return false
            }
        }
    }
  }
  return true  
}