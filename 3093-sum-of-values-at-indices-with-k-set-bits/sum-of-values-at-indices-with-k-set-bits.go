func sumIndicesWithKSetBits(nums []int, k int) int {
    sum := 0
    for i := 0 ; i < len(nums) ; i++ {
        bStr := strconv.FormatInt(int64(i),2)
        s := strings.Replace(bStr,"0","",-1)
        if len(s) == k {
            sum += nums[i]
        }
    }
    return sum
}