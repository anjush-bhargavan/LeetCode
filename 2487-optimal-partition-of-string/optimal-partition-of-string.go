func partitionString(s string) int {
    count := 1
    charMap := make(map[byte]bool)
    for i := 0 ; i < len(s) ; i++ {
        if charMap[s[i]] {
            count++
            clear(charMap)
        }
        charMap[s[i]] = true
    }
    return count
}