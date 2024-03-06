func shipWithinDays(weights []int, days int) int {
    arr := make([]int,len(weights))
    copy(arr,weights)
    sort.Ints(arr)
    n := int(math.Ceil(float64(len(arr))/float64(days)))
    sum := 0
    for i := len(arr)-n ; i < len(arr); i ++  {
        sum += arr[i]
    }
    min,max := arr[len(arr)-1],sum
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