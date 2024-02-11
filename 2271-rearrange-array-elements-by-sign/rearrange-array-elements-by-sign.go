func rearrangeArray(nums []int) []int {
    result := make([]int,len(nums))
    pos,neg := 0,1
    for i := 0 ; i < len(nums); i++{
        if nums[i] < 0 {
            result[neg] = nums[i]
            neg += 2
        }else{
            result[pos] = nums[i]
            pos += 2
        }
    } 
    return result
}