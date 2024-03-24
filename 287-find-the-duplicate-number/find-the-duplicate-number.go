func findDuplicate(nums []int) int {
    i,j := 0,0
    for {
        i = nums[i]
        j = nums[nums[j]]
        if i == j {
            break
        }
    }
    k := 0
    for {
        i = nums[i]
        k = nums[k]
        if i == k {
            return i
        }
    }
    return len(nums)
}