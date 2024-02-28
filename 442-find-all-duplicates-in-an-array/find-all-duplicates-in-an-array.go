func findDuplicates(nums []int) []int {
    result := []int{}
    numMap := make(map[int]bool)

    for i := 0 ; i < len(nums) ; i++ {
        if !numMap[nums[i]] {
            numMap[nums[i]] = true
        }else{
            result = append(result,nums[i])
        }
    }
    
    return result
}