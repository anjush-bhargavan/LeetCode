func findDuplicate(nums []int) int {
   numMap := make(map[int]bool)
    for i := 0 ; i < len(nums) ; i++ {
        if numMap[nums[i]]{
            return nums[i]
        }
        numMap[nums[i]] = true
    }
    return len(nums)
}