func shipWithinDays(weights []int, days int) int {
    sum := 0
    for i := 0 ; i < len(weights); i ++  {
        sum += weights[i]
    }
    min,max := weights[0],sum
    result := 0
    for min <= max {
        mid := (min+max)/2
        if possible(mid,days,weights) {
            result = mid
            max = mid - 1
        }else{
            min = mid + 1
        }
    }
    return result
}

func possible(cap,day int,arr []int) bool {
    weight,i,j := 0,0,0
    for i < len(arr) && j < day {
        if  weight+arr[i] <= cap {
            weight += arr[i]
            i++
        }else{
            j++
            weight = 0
        }
    }
    return i == len(arr)
}