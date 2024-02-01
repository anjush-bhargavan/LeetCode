func sumIndicesWithKSetBits(nums []int, k int) int {
    sum := 0
    for i := 0 ; i < len(nums) ; i++ {
        bStr := strconv.FormatInt(int64(i),2)
        count := 0
        for _,s := range bStr {
           if s == '1'{
               count++
           }
        }
        if count == k {
            sum += nums[i]
        }
    }
    return sum
}