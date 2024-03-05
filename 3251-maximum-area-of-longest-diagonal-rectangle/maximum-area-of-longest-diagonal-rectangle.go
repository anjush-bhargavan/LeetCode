func areaOfMaxDiagonal(dimensions [][]int) int {
    max,area := 0,0
    for i := 0 ; i <len(dimensions) ; i++ {
        diag := (dimensions[i][0]*dimensions[i][0]) + (dimensions[i][1]*dimensions[i][1])
        if diag > max {
            area = dimensions[i][0]*dimensions[i][1]
            max = diag
        }else if diag == max {
                if area < dimensions[i][0]*dimensions[i][1] {
                    area = dimensions[i][0]*dimensions[i][1]
                }
            }
    }
    return area
}