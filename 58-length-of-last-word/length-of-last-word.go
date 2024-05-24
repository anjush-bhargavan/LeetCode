func lengthOfLastWord(s string) int {
    count,flag := 0,false
    for i := len(s)-1 ; i >= 0 ; i-- {
        if s[i] != ' ' {
            count++
            flag = true
        }else if flag {
            break
        }

    }
    return count
}