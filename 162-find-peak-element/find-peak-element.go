func findPeakElement(nums []int) int {
    peak := 0
    if len(nums) == 1 {
        return 0
    }

    for i := 0 ; i < len(nums) ; i++ {
        if i == 0 {
            if nums[i] > nums[i+1] {
                peak = i
                break
            }
        }else if i == len(nums)-1 { 
            if nums[i-1] < nums[i] {
                peak = i
                break
            }
        }else if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
            peak = i
            break
        }
    }
    return peak
}