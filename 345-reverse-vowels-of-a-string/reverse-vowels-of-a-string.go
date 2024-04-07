func reverseVowels(s string) string {
    vowel := ""
    for _,l := range s {
        if IsVowel(l) {
            vowel = string(l) + vowel
        }
    }
    res,j := "",0
    for _,l := range s {
        if IsVowel(l) {
            res += string(vowel[j])
            j++
        }else {
            res += string(l)
        }
    }
    
    return res
}

func IsVowel(s rune) bool {
    return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U'
}