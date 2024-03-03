func findJudge(n int, trust [][]int) int {
    person := make(map[int][]int)
    trusting := make(map[int]bool)
    if n == 1 {
        return n
    }

    for i := 0 ; i < len(trust) ; i++ {
        person[trust[i][1]] = append(person[trust[i][1]],trust[i][0])
        trusting[trust[i][0]] = true
    }
    for k,v := range person {
        if len(v) == n-1 && !trusting[k]{
            return k
        }
    }
    return -1
}