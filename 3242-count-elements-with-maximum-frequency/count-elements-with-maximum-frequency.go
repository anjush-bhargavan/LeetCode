func maxFrequencyElements(nums []int) int {
    numMap := make(map[int]int)
    max,total := 0,0

    for i := 0 ; i < len(nums) ; i++ {
        numMap[nums[i]]++
        v := numMap[nums[i]]
        if max <= v {
            if max == v {
                total += v
            }else{
                max = v
                total = max
            }
        } 
    }
    return total
}