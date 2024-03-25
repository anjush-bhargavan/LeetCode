func findDuplicates(nums []int) []int {
    result := []int{}
    for i := 0 ; i < len(nums) ; i++ {
        if nums[int(math.Abs(float64(nums[i])))-1] > 0 {
            nums[int(math.Abs(float64(nums[i])))-1] *= -1
        }else{
            result = append(result,int(math.Abs(float64(nums[i]))))
    
        }
    }
    return result
}