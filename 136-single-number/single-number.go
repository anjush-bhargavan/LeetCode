func singleNumber(nums []int) int {
    numMap := make(map[int]int)

    for i := 0 ; i < len(nums) ; i ++ {
        numMap[nums[i]]++
    }

    for k,v := range numMap {
        if v == 1 {
            return k
        }
    }
    return nums[0]
}