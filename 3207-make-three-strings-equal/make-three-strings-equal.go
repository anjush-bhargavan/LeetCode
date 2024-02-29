func findMinimumOperations(s1 string, s2 string, s3 string) int {
    count := 0
    for i := 0 ; i < len(s1) && i < len(s2) && i < len(s3) ; i++ {
        if s1[i] == s2[i] && s2[i] == s3[i] {
            count++
        }else{
            break
        }
    }
    if count == 0 {
        return -1
    }
    num := (len(s1) + len(s2) + len(s3)) - (3 * count) 
    return num
}