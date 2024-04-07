func isPowerOfFour(n int) bool {
    pow := 1
    for pow < n {
        pow *= 4
    }
    return pow == n
}