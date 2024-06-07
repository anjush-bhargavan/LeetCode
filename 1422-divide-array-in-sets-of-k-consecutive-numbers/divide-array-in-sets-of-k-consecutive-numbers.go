func isPossibleDivide(nums []int, k int) bool {
    if len(nums) % k != 0 {
        return false
    }
    numMap := make(map[int]int)
    sort.Ints(nums)

    for i := 0 ; i < len(nums) ; i++ {
        numMap[nums[i]]++
    }
    
    for  len(nums) > 0 {
        if numMap[nums[0]] < 1 {
            nums = nums[1:]
        }else{
            count,n := 0,nums[0]
            for count < k {
                if numMap[n] > 0 {
                    numMap[n]--
                    n++
                    count++
                }else{
                    return false
                }
            }
        }

    }

    return len(nums)==0
}