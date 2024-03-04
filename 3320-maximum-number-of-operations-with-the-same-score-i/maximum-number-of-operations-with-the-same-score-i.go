func maxOperations(nums []int) int {
    score,operations := nums[0] + nums[1],1
    for i := 2 ; i < len(nums)-1 ; i += 2 {
        if nums[i] + nums[i+1] == score {
            operations++
        }else{
            break
        }
    }
    return operations
}