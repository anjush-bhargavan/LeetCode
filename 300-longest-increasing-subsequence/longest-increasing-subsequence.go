func lengthOfLIS(nums []int) int {
    dp := make([]int,len(nums))
    for i := 0 ; i < len(nums) ; i++ {
        for j := i-1 ; j >= 0 ; j-- {
           if nums[i] > nums[j] {
            dp[i] = max(dp[i],dp[j] + 1)
           } 
        } 
    }
    max := 1
    for _,s := range dp {
        if max - 1 < s {
            max = s + 1
        }
    }
    return max
}