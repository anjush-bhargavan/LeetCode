func largestTimeFromDigits(arr []int) string {
    time := ""
    numMap := make(map[int]int)
    count := 0

    for _,v := range arr {
        numMap[v]++
        if v < 6 {
            count++
        }
    }
    h := 1
    if count >= 3 {
        h = 2
    }else if count < 2 || (numMap[1] < 1 && numMap[0] < 1){
        return ""
    }
    for i := h ; i >= 0 ; i-- {
        if numMap[i] > 0 {
            time += fmt.Sprintf("%d",i)
            numMap[i]--
            break
        }
    }
    if time == "" {
        return ""
    }

    if time == "2" {
        for i := 3 ; i >= 0 ; i-- {
            if numMap[i] > 0 {
                time += fmt.Sprintf("%d",i)
                numMap[i]--
                break
            }
        }
    }else{
        for i := 9 ; i >= 0 ; i-- {
            if numMap[i] > 0 {
                time += fmt.Sprintf("%d",i)
                numMap[i]--
                break
            }
        }
    }
    if len(time) != 2 {
        return ""
    }
    time += ":"
    for i := 5 ; i >= 0 ; i-- {
        if numMap[i] > 0 {
            time += fmt.Sprintf("%d",i)
            numMap[i]--
            break
        }
    }
    for i := 9 ; i >= 0 ; i-- {
        if numMap[i] > 0 {
            time += fmt.Sprintf("%d",i)
            numMap[i]--
            break
        }
    }
    return time
}

