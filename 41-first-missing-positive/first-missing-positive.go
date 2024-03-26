func firstMissingPositive(nums []int) int {
 
    for i := 0 ; i < len(nums) ; i++ {
        if nums[i] < 0 {
            nums[i] = 0
        }
    }


    for i := 0 ; i < len(nums) ; i++ {
        j := Abs(nums[i]) - 1
        if j >= 0 && j < len(nums) {
            if nums[j] > 0 {
                nums[j] = -1 * nums[j]
            }else if nums[j] == 0 {
                nums[j] = -1 * (len(nums)+1)
            }
        }
    }
    

    for i := 1 ; i < len(nums)+1 ; i++ {
      if nums[i-1] >= 0 {
        return i
      }
    }
    return len(nums) + 1
}


func Abs(a int) int {
    if a < 0 {
        return -1 * a
    }
    return a
}