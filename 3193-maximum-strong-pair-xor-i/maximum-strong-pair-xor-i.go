func maximumStrongPairXor(nums []int) int {
    max := 0

    for i := 0 ; i < len(nums) ; i++ {

        for j := 0 ; j < len(nums) ; j++ {
            if mod(nums[i],nums[j]) <= min(nums[i],nums[j]) {
                if max < nums[i]^nums[j] {
                    max = nums[i]^nums[j]
                }
            }
        }
    }
    return max
}

func mod(x,y int) int {
    if x > y {
        return x-y
    }
    return y-x
}


func min(x,y int) int {
    if x > y {
        return y
    }
    return x
}