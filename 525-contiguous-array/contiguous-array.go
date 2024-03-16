func findMaxLength(nums []int) int {
    difMap := make(map[int]int)
    sum,res := 0,0
    for i := 0 ; i < len(nums) ; i++ {
        if nums[i] == 1 {
            sum++
        }else {
            sum--
        }
        if _,ok := difMap[sum] ; !ok {
            difMap[sum] = i
        }
        if sum == 0 {
            res = max(res,i+1)
        }else{
            idx := difMap[sum]
            res = max(res,i - idx)
        }
    }
    return res
}
