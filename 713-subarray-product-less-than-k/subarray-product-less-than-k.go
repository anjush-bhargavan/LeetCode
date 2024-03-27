func numSubarrayProductLessThanK(nums []int, k int) int {
    count,product,i,j := 0,1,0,0

    for j < len(nums) && i < len(nums) {
        product *= nums[j]
       
        for i < len(nums) && product >= k && product > 1 {
            product /= nums[i]
            i++
        }

        if  product < k {
            count += (j-i) + 1
            j++
        }else{
            i++
        }
    }
    return count
}