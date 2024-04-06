func minRemoveToMakeValid(s string) string {
    stack := []int{}
    indexMap := make(map[int]bool)
    for i := 0 ; i < len(s) ; i++ {
        if s[i] == '(' {
            stack = append(stack,i)
        }else if s[i] == ')' {
            if len(stack) == 0 {
                indexMap[i] = true
            }else{
                stack = stack[:len(stack)-1]
            }
        } 
    }

    for _,i := range stack {
        indexMap[i] = true
    }
    var res strings.Builder 
    for i,l := range s {
        if !indexMap[i] {
            res.WriteRune(l)
        }
    }
    return res.String()
}
