func maxCoins(piles []int) int {
    max,count := 0,0
    sort.Ints(piles)
    for i := len(piles)-2 ; count < len(piles)/3 ; i -= 2 {
        count++
        max +=piles[i]
    } 
    return max
}