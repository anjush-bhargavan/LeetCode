func lastVisitedIntegers(words []string) []int {
    numCount,prevCount := 0,0
    result,last := []int{},[]int{}
    for i := 0 ; i < len(words) ; i++ {
        if words[i] != "prev" {
            num,_ := strconv.Atoi(words[i])
            last = append(last, num)
            numCount++
            prevCount = 0
        }else{
            prevCount++
            idx := numCount - prevCount
              if idx >= 0 {
                result = append(result,last[idx])
            }else{
                result = append(result,-1)
            }
        }
    }
    return result
}