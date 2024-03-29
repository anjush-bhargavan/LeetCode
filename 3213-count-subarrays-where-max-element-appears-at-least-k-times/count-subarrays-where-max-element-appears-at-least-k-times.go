func countSubarrays(nums []int, k int) int64 {
    max := MaxElement(nums)
    countMax,count,i := 0,0,0
    for j :=  range len(nums) {
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



func MaxElement(nums []int) int {
    max := 0
    for _,v := range nums {
        if max < v {
            max = v
        }
    }
    return max
}