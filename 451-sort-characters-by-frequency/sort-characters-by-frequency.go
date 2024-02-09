func frequencySort(s string) string {
    charCount := make(map[byte]int)
    for i := 0 ; i < len(s) ; i++ {
        charCount[s[i]]++
    }
    
    arr := [][]int{}
    for k,v := range charCount {
        arr = append(arr,[]int{int(k),v})
    }

    result := ""
    sortedArr := MergeSort(arr)
    for i := len(sortedArr)-1 ; i >= 0 ; i-- {
       rs := strings.Repeat(string(sortedArr[i][0]),sortedArr[i][1])
       result += rs
    }
   return result
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