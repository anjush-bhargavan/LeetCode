func isPowerOfTwo(n int) bool {
    if n <= 0  {
        return false
    }
    result := 1 
    for result < n {
        result *= 2
    }
    return result == n
}