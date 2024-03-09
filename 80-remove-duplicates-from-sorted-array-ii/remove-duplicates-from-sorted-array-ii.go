func removeDuplicates(nums []int) int {
    i,flag := 0,false
    for i < len(nums)-1 {
        if nums[i] == nums[i+1] {
            if flag {
                nums = append(nums[:i],nums[i+1:]...)
            }else{
                i++
                flag = true
            }
        }else{
            flag = false
            i++
        }
    }
    return len(nums)
}