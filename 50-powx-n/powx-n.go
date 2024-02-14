func myPow(x float64, n int) float64 {
    result,flag,i := x,false,1
    if x == 1.00 || n == 0{
        return 1
    }
    if n < 0 {
        flag = true
        n *= -1
    }
    for 2*i <= n {
        result *= result
        i *= 2

    }
    for i < n {
        result *= x
        i++
    }
    if flag {
        return 1/result
    }
    return result
}