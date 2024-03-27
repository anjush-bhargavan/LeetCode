func numSubarrayProductLessThanK(nums []int, k int) int {
    if k <= 1 {
        return 0
    }
    count,product,i,j := 0,1,0,0

    for j < len(nums) && i < len(nums) {
        product *= nums[j]
       
        for i < len(nums) && product >= k {
            product /= nums[i]
            i++
        }

        if  product < k {
            count += (j-i) + 1
            j++
        }
    }
    return count
}