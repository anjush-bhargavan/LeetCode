func intersection(nums1 []int, nums2 []int) []int {
    numMap := make(map[int]bool)

    for i := 0 ; i < len(nums1) ; i++ {
        numMap[nums1[i]] = true
    }

    result := []int{}
    for i := 0 ; i < len(nums2) ; i++ {
        if numMap[nums2[i]] {
            result = append(result,nums2[i])
            numMap[nums2[i]] = false
        }
    }
    return result
}