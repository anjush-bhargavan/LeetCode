func fillCups(amount []int) int {
    max := minTime(amount)
    seconds := (amount[0]+amount[1])/2 
    a := amount[2]+((amount[0]+amount[1])%2)
    seconds += (a/2)
    seconds += (a%2)

    if max > seconds {
        return max
    }
    return seconds
}


func minTime(amount []int) int {
    max := 0
    if amount[0] < amount[1] {
        max = amount[1]
    }else{
        max = amount[0]
    }
    if amount[2] > max {
        max = amount[2]
    }
    return max
}