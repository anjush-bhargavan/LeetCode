func findJudge(n int, trust [][]int) int {
    person := make(map[int]int)
    trusting := make(map[int]int)

    for i := 0 ; i < len(trust) ; i++ {
        person[trust[i][1]]++
        trusting[trust[i][0]]++
    }
    for i := 1 ; i <= n ; i++ {
        if person[i] == n-1 && trusting[i] == 0 {
            return i
        }
    }
    return -1
}