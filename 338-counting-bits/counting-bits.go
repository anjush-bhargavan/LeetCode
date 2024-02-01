func countBits(n int) []int {
    if n == 0 {
        return []int{0}
    }
    bits := make([]int,n + 1)
    i,j,bin := 2,0,1
    bits[0],bits[1] = 0,1
    for i <= n {
        bin *= 2
        j = i + bin
        for i < j {
            bits[i] = bits[i-bin] + 1
            i++
            if i == n+1 {
                return bits
            }
        }
    }
    return bits
}