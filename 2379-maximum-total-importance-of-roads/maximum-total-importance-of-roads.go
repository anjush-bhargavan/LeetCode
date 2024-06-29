func maximumImportance(n int, roads [][]int) int64 {
    numMap := make(map[int]int)

    for i,_ := range roads {
        numMap[roads[i][0]]++
        numMap[roads[i][1]]++
    }

    result := make([][]int,len(numMap))
    i := 0
    for k,v := range numMap {
        result[i] = make([]int,2)
        result[i][0] = v
        result[i][1] = k
        i++
    }

    sort.Slice(result, func(i, j int) bool {
		return result[i][0] > result[j][0]
	})

    sum := 0
    for i,_ := range result {
        sum += result[i][0] * n
        n--
    }

    return int64(sum)
}