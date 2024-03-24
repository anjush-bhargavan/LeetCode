func convert(s string, numRows int) string {
    matrix := make([][]byte,numRows)
    if len(s) <= numRows || numRows == 1 {
        return s
    }
    n := (len(s)/((2*numRows)-2))+1
    for i,_ := range matrix {
        matrix[i] = make([]byte,(n*(numRows-1)))
    }

    j,k,flag := 0,0,false
    for i,_ := range s {
        if j == len(matrix)-1 {
            flag = true
        }
        if j == 0 {
            flag = false
        }
        matrix[j][k] = s[i]
        if  flag {
            j--
            k++
        }
        if !flag {
            j++
        }
    }

    result := ""
    for i,_ := range matrix {
        for j := 0 ; j < len(matrix[i]) ; j++ {
            if matrix[i][j] != 0 {
                result += string(matrix[i][j])
            }
        }
    }
    return result
}