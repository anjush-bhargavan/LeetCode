func maximumOddBinaryNumber(s string) string {
    n := len(s)
    newS := strings.Replace(s,"0","",-1)
    n -= len(newS)
    newS = newS[:len(newS)-1]
    for i := 0 ; i < n ; i++ {
        newS += "0"
    }
    return newS + "1"
}