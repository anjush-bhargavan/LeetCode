func findMinArrowShots(points [][]int) int {
    arr := MergeSort(points)
    count, min,max := 1,arr[0][0],arr[0][1]
    for i := 1 ; i < len(arr) ; i++ {
        if arr[i][0] <= max {
            if max > arr[i][1] {
                max = arr[i][1]
            }
            if min < arr[i][0] {
                min = arr[i][0]
            }
        }else{
            max = arr[i][1]
            min = arr[i][0]
            count++
        }
    }
    return count 
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
		if arr1[i][0] < arr2[j][0]{
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