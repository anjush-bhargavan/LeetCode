func reverseWords(s string) string {
    result,word,flag := "","",false
    for i := 0 ; i < len(s) ; i++ {
        if s[i] != ' ' && !flag {
            word += string(s[i])
        }

        if flag && s[i] != ' ' && word != ""  {
            result = " " + word + result
            word = string(s[i])
            flag = false 
        }
        
        if s[i] == ' ' {
            flag = true
        }else if word == "" {
            word = string(s[i])
            flag = false
        }
    }
    result = word + result
    return result
}