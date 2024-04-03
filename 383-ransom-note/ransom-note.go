func canConstruct(ransomNote string, magazine string) bool {
    magMap := make(map[byte]int)

    for i := 0 ; i < len(magazine) ; i++ {
        magMap[magazine[i]]++
    }

    for i := 0 ; i < len(ransomNote) ; i++ {
        if magMap[ransomNote[i]] > 0 {
            magMap[ransomNote[i]]--
        }else{
            return false
        }
    }
    return true
}