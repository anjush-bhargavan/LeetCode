func canBeEqual(s1 string, s2 string) bool {
    
    for i := 0 ; i < 2 ; i ++ {
        if s1[i] != s2[i] && s1[i] != s2[i+2] {
            return false
        }
    }
    for i := 2 ; i < 4 ; i ++ {
        if s1[i] != s2[i-2] && s1[i] != s2[i] {
            return false
        }
    }
    return true
}