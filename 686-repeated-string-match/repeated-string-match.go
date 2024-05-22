func repeatedStringMatch(a string, b string) int {
    res,result := "",-1
    n := len(b)/len(a)
    for i := 0 ; i <= n+1 ; i++ {
        res += a
    }
    for i := 0 ; i < len(res) ; i++ { 
        if i <= len(res)-len(b) && b == res[i:i+len(b)] {
            result = (i+len(b)) / len(a)
            if (i+len(b)) % len(a) > 0 {
                result ++
            }
            break
        }
    }
    return result
}