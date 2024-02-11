func searchRange(nums []int, target int) []int {
    low,high,idx := 0 , len(nums)-1 , -1
    for low <= high {
        mid := (low+high)/2
        if nums[mid] == target {
            idx = mid
            break
        }else if target < nums[mid] {
            high = mid - 1
        }else {
            low = mid + 1
        }
    }
    last,first := idx,idx
    if idx == -1 {
        return []int{first,last}
    }

    for {
        if first >= 0 &&  nums[first] == target{
            first--
        }else if  last < len(nums) && nums[last] == target {
            last++
        }else {
            break
        }
    }
    return []int{first+1,last-1}
}