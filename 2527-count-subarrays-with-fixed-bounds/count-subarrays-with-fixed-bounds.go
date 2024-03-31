func countSubarrays(nums []int, minK int, maxK int) int64 {
    minIdx,maxIdx,badIdx,count := -1,-1,-1,0

    for k := range len(nums) {
        if nums[k] < minK || nums[k] > maxK {
            badIdx = k
        }

        if nums[k] == minK {
            minIdx = k
        }
        if maxK == nums[k] {
            maxIdx = k
        } 
        count += max(0,( min(minIdx,maxIdx) - badIdx))
    }
    return int64(count)
}