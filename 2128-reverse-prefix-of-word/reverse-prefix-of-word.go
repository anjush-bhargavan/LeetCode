func reversePrefix(word string, ch byte) string {
    j:=0
    for i:=0;i<len(word);i++{
        if word[i]==ch{
            j=i
            break
        }
    }
    if j==0{
        return word
    }
    result:=""
    for i:=j;i>=0;i--{
        result+=string(word[i])
    }
    result+=word[j+1:]
    return result
}