func pivotInteger(n int) int {
    i,j := 0,n
    sum := (n + 2) * (n - 1) / 2
    for i <= j {
        mid := (i + j) / 2
        pivot := (mid * (mid + 2))
        if pivot == sum {
            return mid + 1
        }else if pivot < sum {
            i = mid + 1
        }else{
            j = mid -1
        }
    } 
    return -1
}