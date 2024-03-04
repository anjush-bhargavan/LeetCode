func bagOfTokensScore(tokens []int, power int) int {
    score := 0
    sort.Ints(tokens)
    i , j,flag := 0,len(tokens)-1, true
    for i <= j && flag {
        flag = false
        if score > 0 && power < tokens[i] {
            power += tokens[j]
            j--
            score--
        }
        if power >= tokens[i] {
            power -= tokens[i]
            score++
            i++
            flag = true
        }
    }
    return score
}