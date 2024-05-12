func jump(nums []int) int {
    arr := make([]int,len(nums))

    for i := 0 ; i < len(nums) ; i++ {
        for j := i+1 ; j <= i+nums[i] && j < len(arr) ; j++ {
          if arr[j] == 0 {
            arr[j] = arr[i]+1
          }else{
            arr[j] = min(arr[j],1+arr[i])
          }  
        } 
    }
    return arr[len(arr)-1]
}