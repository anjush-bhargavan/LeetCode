func searchMatrix(matrix [][]int, target int) bool {
    low,high := 0,len(matrix)-1

    for low <= high {
        mid := (low+high)/2
        if target < matrix[mid][0] {
            high = mid -1
        }else if target > matrix[mid][len(matrix[mid])-1] {
            low = mid + 1
        }else{
            lo,hi := 0,len(matrix[mid])-1
            for lo <= hi {
                mi := (lo+hi)/2
                if target < matrix[mid][mi] {
                    hi = mi - 1
                }else if target > matrix[mid][mi] {
                    lo = mi + 1
                }else{
                    return true
                }
            }
            return false 
        }   
    }
    return false
}