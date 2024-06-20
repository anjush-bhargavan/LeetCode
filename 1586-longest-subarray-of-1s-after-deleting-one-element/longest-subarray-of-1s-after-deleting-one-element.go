func longestSubarray(nums []int) int {
    first,max,count,flag := 0,0,0,false
    for _,n := range nums {
        if n == 1 {
            count++
        }else{
            flag = true
            temp := count
            count -= first 
            first = temp - first
        }
        if max < count {
            max = count
        }
    }
    if !flag {
        max -= 1
    }
    return max
}