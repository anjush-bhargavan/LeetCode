func peakIndexInMountainArray(arr []int) int {
    low , high := 0, len(arr)-1
    for low <= high {
        mid := (low + high)/2
        if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
            return mid
        }else if arr[mid] < arr[mid+1] {
            low = mid
        }else {
            high = mid
        }
    }
    return 0
}