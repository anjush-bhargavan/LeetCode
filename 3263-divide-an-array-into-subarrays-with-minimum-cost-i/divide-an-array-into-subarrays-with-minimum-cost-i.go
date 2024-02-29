func minimumCost(nums []int) int {
    min := nums[0]
    nums = nums[1:]
    sort.Ints(nums)
    return min + nums[0]+nums[1]
}