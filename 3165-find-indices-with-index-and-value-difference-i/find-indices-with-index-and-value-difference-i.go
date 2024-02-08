func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    for i := 0 ; i < len(nums)-indexDifference ; i++ {
        for j := i + indexDifference ; j < len(nums) ; j++ {
            if AbsD(nums[i],nums[j]) >= valueDifference {
                return []int{i,j}
            }
        }
    }
    return []int{-1,-1}
}



func AbsD(a,b int) int {
    if a > b {
        return a-b
    }
    return b-a
}