func divideArray(nums []int, k int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    for i := 0 ; i < len(nums)-2 ; i+=3 {
        if nums[i+2] - nums[i] <= k {
            result = append(result,[]int{nums[i],nums[i+1],nums[i+2]})
        }else{
            return [][]int{}
        } 
    }
    return result
}