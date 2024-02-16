func findLeastNumOfUniqueInts(arr []int, k int) int {
    intMap := make(map[int]int)

    for i := 0 ; i < len(arr) ; i++ {
        intMap[arr[i]]++
    }
    nums := [][]int{}
    for k,v := range intMap {
        temp := []int{k,v}
        nums = append(nums,temp)
    }
    sortedArr := MergeSort(nums)
    j,count := 0,0
    fmt.Println(sortedArr)
    for i := 0 ; i < len(sortedArr) ; i++ {
        if j < k {
            j += sortedArr[i][1]
        }else{
            break
        }
        count++
    }
    if j > k {
        return len(intMap) - count + 1
    }
    return len(intMap) - count

}


func MergeSort(arr [][]int) [][]int {
    if len(arr) == 1 {
        return arr
    }
    mid := len(arr)/2
    left := MergeSort(arr[:mid])
    right := MergeSort(arr[mid:])

    return merge(left,right)
}

func merge(arr1,arr2 [][]int) [][]int {
	combined:=[][]int{}
	i,j:=0,0
	for i<len(arr1) && j <len(arr2){
		if arr1[i][1] < arr2[j][1]{
            combined = append(combined,arr1[i])
            i++
		}else{
			combined = append(combined,arr2[j])
            j++
		}
	}
	for i < len(arr1){
	    combined = append(combined,arr1[i])
        i++
	}
	for j < len(arr2){
		combined = append(combined,arr2[j])
        j++
	}
	return combined
}