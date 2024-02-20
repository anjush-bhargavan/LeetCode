func countVowelStrings(n int) int {
    d4,d3,d2,d1 := 5,10,10,5
    for i := 1 ; i < n ; i++ {
        d4 += d3 
        d3 += d2
        d2 += d1
        d1++

    }
    return d4
}