func numberOfPoints(nums [][]int) int {
    pointMap := make(map[int]bool)
    count := 0
    for i := 0 ; i < len(nums) ; i++ {
        for j := nums[i][0] ; j <= nums[i][1] ; j++ {
            if !pointMap[j] {
                pointMap[j] = true
                count++
            }
        }
    }
    return count
}