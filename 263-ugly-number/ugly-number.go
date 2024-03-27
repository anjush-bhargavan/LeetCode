func isUgly(n int) bool {
    if n <= 0 {
        return false
    }

    arr := []int{2,3,5} 
    for _,v := range arr {
        for n % v == 0 {
            n /=v
        }
    }
    return n == 1
}