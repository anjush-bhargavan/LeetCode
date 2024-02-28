func twoEggDrop(n int) int {
    sum,num := 0,0
    for i := 1 ; sum < n ; i++ {
        sum += i
        num = i
    }
    return num
}