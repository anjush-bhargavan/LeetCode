func minDifference(nums []int) int {
    if len(nums) < 5 {
        return 0
    }
    sort.Ints(nums)
    return min(min((nums[len(nums)-3] - nums[1]),(nums[len(nums)-2] - nums[2])),min((nums[len(nums)-4] - nums[0]),(nums[len(nums)-1] - nums[3])))
}