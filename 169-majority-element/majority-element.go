func majorityElement(nums []int) int {
    numMap := make(map[int]int)
    max,majElement := 0,0
    for i := 0 ; i < len(nums) ; i++ {
        numMap[nums[i]]++
        if max < numMap[nums[i]] {
            max = numMap[nums[i]]
            majElement = nums[i]
        }
    }
    return majElement
}