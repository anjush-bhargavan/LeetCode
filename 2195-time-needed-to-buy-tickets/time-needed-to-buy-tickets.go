func timeRequiredToBuy(tickets []int, k int) int {
    count := 0
    target := tickets[k]
    for i,v := range tickets {
        if i <= k {
            count += min(v,target)
        }else {
            count += min(v,target-1)
        }
    }
    return count
}