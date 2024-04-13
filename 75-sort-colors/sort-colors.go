func sortColors(nums []int)  {
 countOne,countZero,countTwo := 0,0,0
 for _,n := range nums {
    if n == 0 {
        countZero++
    }else if n == 1 {
        countOne++
    }else{
        countTwo++
    }
 }
 for i,_ := range nums {
    if countZero > 0 {
        nums[i] = 0
        countZero--
    }else if countOne > 0 {
        nums[i] = 1 
        countOne--
    }else if countTwo > 0 {
        nums[i] = 2
        countTwo--
    }
 }
}