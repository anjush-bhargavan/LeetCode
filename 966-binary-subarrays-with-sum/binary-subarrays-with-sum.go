func numSubarraysWithSum(nums []int, goal int) int {
    count := 0
    for i := 0 ; i < len(nums) ; i++ {
        sum,j := 0,i
        for sum <= goal && j < len(nums)  {
            sum += nums[j]
            j++
            if sum == goal {
                count++
            }
        }
    }
 
    return count
}

