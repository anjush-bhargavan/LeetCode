func insert(intervals [][]int, newInterval []int) [][]int {
    result := [][]int{}
    temp := make([]int,2)
    flag := false
    if len(intervals) == 0 || intervals[len(intervals)-1][1] < newInterval[0] {
        result = append(intervals,newInterval) 
        return result
    }else if intervals[0][0] > newInterval[1] {
        result = append(result,newInterval)
        result = append(result,intervals...) 
        return result
    }

    for i := 0 ; i < len(intervals) ; i++ {
        if intervals[i][1] >= newInterval[0] && !flag {
            if intervals[i][0] >= newInterval[0] {
                if intervals[i][0] <= newInterval[1] {
                    temp[0] = newInterval[0]
                }else{
                    result = append(result,newInterval)
                    result = append(result,intervals[i:]...)
                    return result
                }
            }else{
                temp[0] = intervals[i][0]
            }
            flag = true
            if  intervals[i][1] >= newInterval[1] {
                temp[1] = intervals[i][1]
                result = append(result,temp)
                result = append(result,intervals[i+1:]...)
                return result
            }else {
                temp[1] = newInterval[1]
            }
            continue
        }
        if !flag {
            result = append(result,intervals[i])
        }else{
            if  intervals[i][0] <= temp[1] {
                if intervals[i][1] <= temp[1] {
                    continue
                }
                temp[1] = intervals[i][1]
                result = append(result,temp)
                result = append(result,intervals[i+1:]...)
                flag = false
                break
            }else{
                result = append(result,temp)
                result = append(result,intervals[i:]...)
                flag = false
                break
            }
        }
    }
    if flag {
        result = append(result,temp)
    }
    return result
}
