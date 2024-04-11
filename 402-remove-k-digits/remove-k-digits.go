func removeKdigits(num string, k int) string {
    stack := []rune{}
    
    for _,s := range num {
        for len(stack) > 0 && stack[len(stack)-1] > s && k > 0 {
            stack = stack[:len(stack)-1]
            k--
        }
        stack = append(stack, s)
    }
    stack = stack[:len(stack)-k]
    
    res := ""
    for _,s := range stack {
        res += string(s)

    }

    n, ok := new(big.Int).SetString(res, 10)
    if !ok {
        return "0"
    }
    res = fmt.Sprintf("%v", n)
    return res
}