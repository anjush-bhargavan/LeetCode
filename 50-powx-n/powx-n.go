func myPow(x float64, n int) float64 {
    if n < 0 {
        return 1/pow(x,-n)
    }
 
    return pow(x,n)
}

func pow(x float64,n int) float64 {
    if n == 0 {
        return 1
    }
    result := pow(x,n/2)
    result *= result
    if n%2 == 1 {
        result *= x
    }
    return result
}