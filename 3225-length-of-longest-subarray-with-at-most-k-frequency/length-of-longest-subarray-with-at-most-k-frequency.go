func maxSubarrayLength(nums []int, k int) int {
    numMap := make(map[int]int)
    length,i,j := 0,0,0
    for j < len(nums) {
        numMap[nums[j]]++
        if numMap[nums[j]] <= k {
            if length < (j-i) + 1 {
                length = j - i + 1 
            } 
        }else{
            for numMap[nums[j]] > k {
                numMap[nums[i]]--
                i++
            }
        }
        j++
    }
    return length
}