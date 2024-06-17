func judgeSquareSum(c int) bool {
    start,sum := 0,0
    end := int(math.Sqrt(float64(c)))
    for start <= end {
        sum = (start*start) + (end*end)
        if sum < c {
            start++
        }else if sum > c {
            end--
        }else{
            break
        }
    }
    return sum == c
}