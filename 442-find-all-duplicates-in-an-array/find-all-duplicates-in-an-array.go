func findDuplicates(nums []int) []int {
    result := []int{}
    for i := 0 ; i < len(nums) ; i++ {
        if nums[Abs(nums[i])-1] < 0 {
            result = append(result,Abs(nums[i]))
        }
        nums[Abs(nums[i])-1] *= -1
    }
    return result
}


func Abs(x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}