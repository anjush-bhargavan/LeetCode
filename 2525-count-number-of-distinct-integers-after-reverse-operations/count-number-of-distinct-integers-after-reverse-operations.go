func countDistinctIntegers(nums []int) int {
    intMap := make(map[int]bool)
    for i := 0 ; i < len(nums) ; i++ {
        intMap[nums[i]] = true
        temp,rev := nums[i],0
        for temp > 0 {
            rev = rev*10 + (temp % 10)
            temp /= 10
        }
        intMap[rev] = true 
    }
    return len(intMap)
}