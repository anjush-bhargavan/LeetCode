func minSubArrayLen(target int, nums []int) int {
    length,sum,i,j,flag := len(nums),0,0,0,false

    for i < len(nums) {
       
        if sum >= target {
            if length >= (j-i) {
            length = j - i 
            flag = true
            }
        }
        if j < len(nums) && sum <= target {
            sum += nums[j]
            j++
        }else{
            sum -= nums[i]
            i++
        }
    }
    if !flag {
        return 0
    }
    return length
}