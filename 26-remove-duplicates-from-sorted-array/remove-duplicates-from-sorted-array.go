func removeDuplicates(nums []int) int {
    numMap := make(map[int]bool)
    idx := 0
    for i := 0 ; i < len(nums) ; i++ {
        if !numMap[nums[i]] {
            nums[idx] = nums[i]
            idx++
            numMap[nums[i]] = true
        }
    }
    return idx
}