func isIsomorphic(s string, t string) bool {
    charMapS := make(map[byte]byte)
    charMapT := make(map[byte]byte)

    for i := 0 ; i < len(s) ; i++ {
        if charMapS[s[i]] == 0  && charMapT[t[i]] == 0 {
            charMapS[s[i]] = t[i]
            charMapT[t[i]] = s[i]
        }else if charMapS[s[i]] != t[i] || charMapT[t[i]] != s[i] {
            return false
        }
    }
    return true
}