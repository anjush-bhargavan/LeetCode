func numTrees(n int) int {
    var num,i float64
    num = 1
    
    for i = 2 ; i <= float64(n) ; i++ {
        num *= ((float64(n) + i) / i)
    }
    result := math.Round(num)

    return int(result)  
} 
