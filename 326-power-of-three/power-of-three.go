func isPowerOfThree(n int) bool {
    result := 1
    for result < n {
        result *= 3
    }
    return result == n
}