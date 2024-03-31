func sumOfTheDigitsOfHarshadNumber(x int) int {
    sum,n := 0,x

    for n > 0 {
        sum += (n % 10)
        n /= 10
    }
    if x % sum == 0 {
        return sum
    }
    return -1
}