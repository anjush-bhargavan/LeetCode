func sequentialDigits(low int, high int) []int {
    lowS := strconv.Itoa(low)
    highS := strconv.Itoa(high)
    result := []int{}
    seq := "123456789"
    for i := len(lowS) ; i <= len(highS) ; i++{
        
        for j := 0 ; j <= len(seq)-i ; j++ {
            seqN,_ := strconv.Atoi(seq[j:j+i])
            if seqN >= low && seqN <= high {
                result = append(result,seqN)
            }
        } 
    }
    return result
}