func maximumOddBinaryNumber(s string) string {
    n := len(s)
    countOne := 0
    for _,l := range s {
        if l == '1' {
            countOne++
        }
    }
    newS := ""
    for i := 1 ; i < n ; i++ {
        if countOne > 1 {
            newS += "1"
            countOne--
        }else{
            newS += "0" 
        }
    }
    return newS + "1"
}