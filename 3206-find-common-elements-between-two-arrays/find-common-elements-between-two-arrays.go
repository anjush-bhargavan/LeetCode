func findIntersectionValues(nums1 []int, nums2 []int) []int {
    num1Map := make(map[int]bool)
    num2Map := make(map[int]bool)

    for i := 0 ; i < len(nums1) ; i++ {
        num1Map[nums1[i]] = true
    }
    count1,count2 := 0,0
    for i := 0 ; i < len(nums2) ; i++ {
        if num1Map[nums2[i]] {
            count1++
        }
        num2Map[nums2[i]] = true
    }
    for i := 0 ; i < len(nums1) ; i++ {
        if num2Map[nums1[i]] {
            count2++
        }
    }
    return []int{count2,count1}
}