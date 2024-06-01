func countSubstrings(s string) int {
    count := 0

    for i := range s {
        l,r := i,i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            count++
            l--
            r++
        }
        l,r = i,i+1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            count++
            l--
            r++
        }
    }
    return count
}