func maxSubArray(nums []int) int {
    i,max,sum := 1,nums[0],nums[0]

    for i < len(nums) {
        if sum < 0 {
            sum = 0
        }
        sum += nums[i]
        i++
        if max < sum {
            max = sum
        }
    }
    return max
}