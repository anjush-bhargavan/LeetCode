func minOperations(boxes string) []int {
    result := make([]int,len(boxes))
    count := 0
    
    for i := 1 ; i < len(result); i++ {
        if boxes[i-1] == '1' {
            count++
        }
        result[i] = result[i-1] + count
    }
    
    result1 := make([]int,len(boxes))
    count = 0
     for i := len(result)-2 ; i >= 0; i-- {
        if boxes[i+1] == '1' {
            count++
        }
        result1[i] += result1[i+1] + count
    }
    for i,v := range result1 {
        result[i] += v
    }
    return result
}