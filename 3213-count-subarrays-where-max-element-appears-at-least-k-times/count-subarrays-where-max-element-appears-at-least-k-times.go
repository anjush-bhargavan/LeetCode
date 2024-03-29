func countSubarrays(nums []int, k int) int64 {
    countMax,count,i,max  := 0,0,0,0
    for j :=  range len(nums) {

        if max < nums[j] {
            max = nums[j]
            countMax = 0
            count = 0
        }

        if nums[j] == max {
            countMax++
        }

        for countMax > k || (i <= j && countMax == k && nums[i] != max){
            if nums[i] == max {
                countMax--
            }
            i++
        }

        if countMax == k {
            count += i + 1
        }
    }
    return int64(count)
}
