func customSortString(order string, s string) string {
    orderMap := make(map[string]int)

    for i := 0 ; i < len(order) ; i++ {
        orderMap[string(order[i])] = i+1
    }
    s1 := strings.Split(s,"")
    // k := len(s1)-1
    for i := 0 ; i < len(s1) ; i++ {
        minIndex := i
        for j := i+1 ; j < len(s1) ; j++ {
            if _,ok := orderMap[s1[j]] ; ok {
                if orderMap[s1[j]] < orderMap[s1[minIndex]] {
                    minIndex = j
                }
            }
        }
        if minIndex != i {
            s1[i],s1[minIndex] = s1[minIndex],s1[i]
        }
    }
    result := strings.Join(s1,"")
    return result
}