func productExceptSelf(nums []int) []int {
    result := make([]int,len(nums))
    product,count,idx := 1,0,0
    for i := 0 ; i < len(nums) ; i++ {
        if nums[i] == 0 {
            count++
            idx = i
        }else{
            product *= nums[i]
        }
    }
    if count > 1 {
        return result
    }else if count == 1 {
        result[idx] = product
        return result
    }
    for i := 0 ; i < len(nums) ; i++ {
        result[i] = product/nums[i]
    }
    return result
}