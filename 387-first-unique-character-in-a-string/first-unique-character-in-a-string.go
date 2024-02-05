func firstUniqChar(s string) int {
    charMap := make(map[byte]int)
    for i := 0 ; i < len(s) ; i++ {
        charMap[s[i]]++
    } 
    for i := 0 ; i < len(s) ; i++ {
        if charMap[s[i]] == 1 {
            return i
        }
    }
    return -1
}