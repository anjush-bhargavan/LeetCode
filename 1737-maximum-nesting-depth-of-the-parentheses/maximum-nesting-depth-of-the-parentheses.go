func maxDepth(s string) int {
    count,max := 0 ,0 
    for i := 0 ; i < len(s) ; i++ {
        if s[i] == '(' {
            count++
        }
        if s[i] == ')' {
            count--
        }
        if max < count {
            max = count
        }
    }
    return max
}