func countKeyChanges(s string) int {
    word := strings.ToLower(s)
    count := 0
    for i := 1 ; i < len(word) ; i++ {
        if word[i] != word[i-1] {
            count++
        }
    }
    return count
}