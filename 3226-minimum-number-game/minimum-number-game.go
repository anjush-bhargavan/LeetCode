func numberGame(nums []int) []int {
    arr := make([]int,len(nums))
    sort.Ints(nums)

    for i := 0 ; i <len(nums) ; i+=2 {
        arr[i] = nums[i+1]
        arr[i+1] = nums[i]
    }
    return arr
}