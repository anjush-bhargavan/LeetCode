func subsets(nums []int) [][]int {
    result := [][]int{}
    num := []int{}
    add(&result,nums,&num,0)
    return result
}


func add(arr *[][]int,nums []int,num *[]int,i int) {
    if i >= len(nums) {
        arr1 := make([]int,len(*num))
        copy(arr1,(*num))
        (*arr) = append(*arr,arr1)
        return
    }
    (*num) = append(*num,nums[i])
    add(arr,nums,num,i+1)
    (*num) = (*num)[:len(*num)-1]
    add(arr,nums,num,i+1)
} 