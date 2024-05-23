func beautifulSubsets(nums []int, k int) int {
    count := make(map[int]int)
    result := SubCount(nums,&count,0,k)
    return result-1
}


func SubCount(nums []int,count *map[int]int,i,k int) int {
    if len(nums) == i {
        return 1
    }

    res := SubCount(nums,count,i+1,k)
    if (*count)[nums[i]-k] <= 0 && (*count)[k+nums[i]] <= 0 {
        (*count)[nums[i]]++
        res += SubCount(nums,count,i+1,k)
        (*count)[nums[i]]--
    }
    return res
}
