func makeGood(s string) string {
    s1,i := s,0
    for i < len(s1)-1 {
        if absDiff(s1[i],s1[i+1]) == 32 {
            s1 = s1[:i] + s1[i+2:]
            i = 0
        }else{
            i++
        }
        
    }
    return s1
}

func absDiff(a,b byte) byte {
    if a > b {
        return a - b
    }
    return b - a
}