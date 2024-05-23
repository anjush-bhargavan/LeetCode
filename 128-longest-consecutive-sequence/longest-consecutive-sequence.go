func longestConsecutive(nums []int) int {
    numMap := make(map[int]bool)
    for i := 0 ; i < len(nums) ; i++ {
        numMap[nums[i]] = true
    }

    start := []int{}
    for k,_ := range numMap {
        if !numMap[k-1] {
            start = append(start,k)
        }
    }

    count,max := 0,0
    for _,v := range start {
        i := v
        for numMap[i] {
            count++
            i++
        } 
        if max < count {
            max = count
        }
        count = 0
    }

    return max
}
