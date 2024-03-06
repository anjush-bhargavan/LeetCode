func maxProfit(prices []int) int {
    min,max,profit,minIdx := prices[0],0,0,0

    for i := 1 ; i < len(prices) ; i++ {
        if min > prices[i] {
            min = prices[i]
            minIdx = i
            max = 0
        }
        if max < prices[i] && i > minIdx{
            max = prices[i]
        }
        if max - min > profit {
            profit = max - min
        }
    }
    return profit
}