func removeStars(s string) string {
    result := []rune{}
    for i := 0 ; i < len(s) ; i++ {
        if s[i] == '*' {
            result = result[:len(result)-1]
        }else{
            result = append(result,rune(s[i]))
        }
    }
    return string(result)  
}


