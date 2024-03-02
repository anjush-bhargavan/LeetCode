func returnToBoundaryCount(nums []int) int {
    position,count := 0,0
    for i := 0 ; i < len(nums) ; i++ {
        position += nums[i]
        if position == 0 {
            count++
        }
    }
    return count
}