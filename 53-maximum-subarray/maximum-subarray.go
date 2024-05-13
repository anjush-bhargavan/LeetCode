func maxSubArray(nums []int) int {
    sum,max := nums[0],nums[0]
    for i := 1 ; i < len(nums) ; i++ {
        if sum < 0 {
            sum = 0
        }
        sum += nums[i]
        if max < sum {
            max = sum
        }
    }
    return max
}