func largestPerimeter(nums []int) int64 {
    sort.Ints(nums)

    arr := make([]int,len(nums))
    sum := -1
    arr[0] = nums[0]
    for i := 1 ; i < len(arr) ; i++ {
        arr[i] = nums[i]+ arr[i-1]
    }

    for i := len(arr)-1 ; i > 0 ; i-- {
        if arr[i-1] > nums[i] {
            sum = arr[i-1] + nums[i]
            break
        }
    }
    return int64(sum)
}











    // // sum := nums[0] + nums[1]
    // temp,count,flag := 0,2,false
    // for i := 2 ; i < len(nums) ; i++ {
    //     if sum >= nums[i] {
    //         flag = false
    //         if sum == nums[i] {
    //             flag = true
    //         }
    //         sum += nums[i]
    //         count++
    //         temp = nums[i]
    //     }
    
    // }
    // if flag {
    //     sum -= temp
    //     count--
    // }
    // if count == 2 {
    //     return -1
    // }