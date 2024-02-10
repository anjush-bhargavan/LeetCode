func searchRange(nums []int, target int) []int {
    first,last,flag := -1,-1,false
    for i := 0 ; i <len(nums) ; i++ {
        if nums[i] == target && !flag{
            first = i
            flag = true
        }
        if nums[i] == target && flag{
            last = i
            flag = true
        }
    }
    return []int{first,last}
}