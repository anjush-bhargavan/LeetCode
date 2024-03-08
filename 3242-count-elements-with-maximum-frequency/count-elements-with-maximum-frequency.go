func maxFrequencyElements(nums []int) int {
    numMap := make(map[int]int)
    max,total := 0,0

    for i := 0 ; i < len(nums) ; i++ {
        numMap[nums[i]]++
    }

    for _,v := range numMap {
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