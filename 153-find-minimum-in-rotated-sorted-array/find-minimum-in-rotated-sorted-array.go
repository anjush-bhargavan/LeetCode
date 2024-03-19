func findMin(nums []int) int {
    low,mid,high := 0,0,len(nums)-1
    for low < high {
        mid = (low+high)/2
        if nums[mid] > nums[mid+1] {
            return nums[mid+1]
        }else if nums[mid] < nums[low] {
            high = mid 
        }else{
            low = mid + 1
        }
    }

    return nums[0]
}