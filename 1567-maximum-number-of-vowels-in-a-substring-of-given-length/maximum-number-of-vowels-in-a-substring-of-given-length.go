func maxVowels(s string, k int) int {
    j,max,count := 0,0,0
    for i := range s {
        for j < len(s) && j-i < k {
            if IsVowel(s[j]) {
                count++
            }
            j++
        }
        if max < count {
            max = count
        }
        if IsVowel(s[i]) {
            count--
        }
    }
    return max
}


func IsVowel(s byte) bool {
    return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' 
}