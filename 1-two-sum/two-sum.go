func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i,n := range nums {
        numMap[n] = i
    }
    for i,n := range nums {
        if v,ok := numMap[target-n] ; ok && v != i {
            return []int{v,i}
        }
    }
    return []int{nums[0],nums[1]}
}