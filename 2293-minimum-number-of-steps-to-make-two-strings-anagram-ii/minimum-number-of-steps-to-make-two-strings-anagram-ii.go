func minSteps(s string, t string) int {
    sMap := make(map[byte]int)
    tMap := make(map[byte]int)
    for i := 0 ; i < len(s) ; i++ {
        sMap[s[i]]++
    }
    for i := 0 ; i < len(t) ; i++ {
        tMap[t[i]]++
    }
    count := 0
    for i := 0 ; i < len(s) ; i++ {
        if tMap[s[i]] > 0 {
            tMap[s[i]]--

        }else{
            count++
        }
    } 
    for i := 0 ; i < len(t) ; i++ {
        if sMap[t[i]] > 0 {
            sMap[t[i]]--

        }else{
            count++
        }
    }
    return count
}