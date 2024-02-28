func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)

    for i := 0 ; i < len(costs) ; i++ {
        if coins - costs[i] >= 0 {
            coins -= costs[i]
        }else {
            return i
        }
    }
    return len(costs)
}