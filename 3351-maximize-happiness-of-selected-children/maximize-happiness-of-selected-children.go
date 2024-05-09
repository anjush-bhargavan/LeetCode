func maximumHappinessSum(happiness []int, k int) int64 {
    sort.Ints(happiness)
    sum := 0
    for i,j,l := len(happiness)-1,k,0 ; j > 0 ; i-- {
        if happiness[i] >= l {
            sum += happiness[i]
            sum -= l
        }
        l++
        j--
    }
    return int64(sum)
}