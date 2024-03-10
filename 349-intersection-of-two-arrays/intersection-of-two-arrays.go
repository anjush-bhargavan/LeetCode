func intersection(nums1 []int, nums2 []int) []int {
    result := []int{}
    numMap := make(map[int]bool)
    sort.Ints(nums1)
    sort.Ints(nums2)
    i,j := 0,0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            i++
        }else if nums1[i] > nums2[j] {
            j++
        }else{
            if !numMap[nums2[j]] {
                result = append(result,nums2[j])
                numMap[nums2[j]] = true
            }
            i++
            j++
        }
    }
    return result
}