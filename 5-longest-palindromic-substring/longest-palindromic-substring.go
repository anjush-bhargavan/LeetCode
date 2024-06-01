func longestPalindrome(s string) string {
    res,max := "",0
    for i := range s {
        l,r := i,i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if max <= (r-l) {
                max = r-l
                res = s[l:r+1]
            }
            l--
            r++
        }
        l,r = i,i+1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if max < (r-l) {
                max = r-l
                res = s[l:r+1]
            }
            l--
            r++
        }
    }
    return res
}