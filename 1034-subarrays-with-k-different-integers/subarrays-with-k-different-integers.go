func subarraysWithKDistinct(nums []int, k int) int {
    count,i,j := 0,0,0
    numMap := make(map[int]int)

    for l := range len(nums) {

        numMap[nums[l]]++
    
            for len(numMap) > k {
                numMap[nums[j]]--
                if numMap[nums[j]] == 0 {
                    delete(numMap,nums[j])
                }
                j++
                i = j
            }

            for numMap[nums[j]] > 1 {
                numMap[nums[j]]--
                j++
            }

        if len(numMap) == k {
            count += (j-i) + 1
        }

    }
    return count
}
