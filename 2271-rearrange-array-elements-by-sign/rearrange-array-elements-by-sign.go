func rearrangeArray(nums []int) []int {
    negArray,posArray,result := []int{},[]int{},[]int{}
    for i := 0 ; i < len(nums) ; i++ {
        if nums[i] < 0 {
            negArray = append(negArray,nums[i])
        }else{
            posArray = append(posArray,nums[i])
        }
    }
    for i := 0 ; i < len(posArray) ; i++ {
        result = append(result,posArray[i])
        result = append(result,negArray[i])
    }
    return result
}