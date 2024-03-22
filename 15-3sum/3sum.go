func threeSum(nums []int) [][]int {
    result := [][]int{}
    sort.Ints(nums)
    for i := 0 ; i < len(nums) ; i++ {

        j :=  i + 1
        k := len(nums) - 1

        for i < len(nums)-1 && nums[i] == nums[i+1] {
            i++
        }
        
        for j < k {
            v :=  nums[i] + nums[j] + nums[k]  
            if v == 0  {
            
                temp := []int{nums[i],nums[j],nums[k]}
                result = append(result,temp)

                for j < k && nums[j] == nums[j+1] {
                    j++
                }
                for j < k && nums[k] == nums[k-1] {
                    k--
                }
                j++
                k--
            }else if v < 0 {
                j++
            }else{
                k-- 
            }
        }

    }
    return result  
}
