func minimumSum(nums []int) int {
    minSum := -1
    for i := 0 ; i < len(nums)-2 ; i++ {
        for j := i+1 ; j < len(nums)-1 ; j++ {
            for k := j+1 ; k < len(nums) ; k++ {
                if nums[i] < nums[j] && nums[j] > nums[k] {
                    sum := nums[i]+nums[j]+nums[k]
                    if minSum > sum || minSum == -1{
                        minSum = sum
                    } 
                }
            }
        }
    }
    return minSum
}