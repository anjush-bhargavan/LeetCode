func findMaxK(nums []int) int {
    result := -1
    m := make(map[int]bool)
    for _,n := range nums {
        m[n]=true
    }
    for _,n := range nums {
        if m[-n]{
            if result < n {
                result = n
            }
        }
    }
    return result
}