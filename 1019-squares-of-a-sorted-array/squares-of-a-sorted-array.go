func sortedSquares(nums []int) []int {
    result := make([]int,len(nums))
    i,j,k := 0,len(nums)-1,len(nums)-1

    for k = len(nums)-1 ; k >= 0 ; k-- {
        if nums[i] < 0 {
            nums[i] *= -1
        }

        if nums[i] >  nums[j]{
            result[k] = nums[i] * nums[i]
            i++ 
        }else{
            result[k] = nums[j] * nums[j]
            j--
        }
    }
    return result
}