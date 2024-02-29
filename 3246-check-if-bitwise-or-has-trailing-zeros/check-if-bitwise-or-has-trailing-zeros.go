func hasTrailingZeros(nums []int) bool {
    count := 0
    for i := 0 ; i < len(nums) ; i++ {
        if nums[i] % 2 == 0 {
            count++
        }
        if count >= 2 {
            return true
        }
    }
    return false
}