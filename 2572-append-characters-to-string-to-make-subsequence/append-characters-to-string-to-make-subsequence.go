func appendCharacters(s string, t string) int {
    i,j,count := 0,0,0

    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i++
            j++
            count++
        }else{
            i++
        } 
    }
    return len(t) - count
}