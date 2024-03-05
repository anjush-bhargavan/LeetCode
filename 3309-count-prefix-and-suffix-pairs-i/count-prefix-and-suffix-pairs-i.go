func countPrefixSuffixPairs(words []string) int {
    count := 0

    for i := 0 ; i < len(words) ; i++ {
    
        for j := i+1 ; j < len(words) ; j++ {
            if isPrefixAndSuffix(words[i],words[j]) {
                count++
            }
        }
    }
    return count
}


func isPrefixAndSuffix(str1,str2 string) bool {
    return len(str2) >= len(str1) && str2[:len(str1)] == str1 && str2[len(str2)-len(str1):] == str1
}