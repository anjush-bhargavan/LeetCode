func findMinArrowShots(points [][]int) int {
    Sort2D(points)
    count,max := 1,points[0][1]
    for i := 1 ; i < len(points) ; i++ {
        if points[i][0] <= max {
            if max > points[i][1] {
                max = points[i][1]
            }
        }else{
            max = points[i][1]
            count++
        }
    }
    return count 
}

func Sort2D(arr [][]int) {
    sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})

}