func canBeEqual(s1 string, s2 string) bool { 
    for i := 0 ; i < 2 ; i ++ {
        s := string(s1[i]) + string(s1[i+2])
        if s != string(s2[i])+string(s2[i+2]) && s != string(s2[i+2])+string(s2[i]) {
            return false
        }
    }
    return true
}