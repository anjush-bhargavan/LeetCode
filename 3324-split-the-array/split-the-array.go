func isPossibleToSplit(nums []int) bool {
    numMap := make(map[int]int)

    for _,v := range nums {
        numMap[v]++
        if numMap[v] > 2 {
            return false
        }
    }
    return true
}