func topKFrequent(nums []int, k int) []int {
   countMap := make(map[int]int)
   for  i := 0 ; i < len(nums) ; i++ {
       countMap[nums[i]]++
   }

    arr :=[][]int{}
    for i,v := range countMap {
        arr = append(arr,[]int{i,v}) 
    }

    sortedArr := MergeSort(arr)
    output :=[]int{}
    j := len(sortedArr)-1
    for i := 0 ; i < k ; i++ {
        output = append(output,sortedArr[j][0])
        j--
    }
    return output
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