func timeRequiredToBuy(tickets []int, k int) int {
    count := 0
    personMap := make(map[int]int)
    for i := 0 ; i < len(tickets) ; i++ {
        personMap[i] = tickets[i]
    }
    i := 0
    for personMap[k] > 0 {
        if personMap[i] > 0 {
            personMap[i]--
            count++
        }
        i++
        if i == len(tickets) {
            i = 0
        }
    }
    return count
}