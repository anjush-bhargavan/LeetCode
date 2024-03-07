func removeElement(nums []int, val int) int {

    for i := 0 ; i  < len(nums) ; i++ {
        if val == nums[i] {
            nums = append(nums[:i],nums[i+1:]...) 
            i--
        }
    }
    return len(nums)
}