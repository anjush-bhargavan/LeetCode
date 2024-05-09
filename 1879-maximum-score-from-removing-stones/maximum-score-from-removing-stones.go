func maximumScore(a int, b int, c int) int {
    if a == 0 && b == 0 || b == 0 && c == 0 || a == 0 && c == 0 {
        return 0
    }
    if a < b {
        b--
        if a < c {
            c--
        }else{
            a--
        }
    }else{
        a--
        if b < c {
            c--
        }else{
            b--
        }
    }
    return maximumScore(a,b,c) + 1
}