func minOperations(nums []int) int {
    count,i := 0,0

    for i < len(nums) {
        if nums[i] == 1 {
            i++
        }else{
            nums[i] = 1
            if i+1 < len(nums) {
                if nums[i+1] == 1 {
                    nums[i+1] = 0
                }else{
                    nums[i+1] = 1
                }
            }else{
                return -1
            }
            if i+2 < len(nums) {
                if nums[i+2] == 1 {
                    nums[i+2] = 0
                }else{
                    nums[i+2] = 1
                }
            }else{
                return -1
            }
            count++
        }
    }
    return count
}