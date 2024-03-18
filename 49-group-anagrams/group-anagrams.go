func groupAnagrams(strs []string) [][]string {

    charMap := make(map[string][]string)
    result := [][]string{}

    for i := 0 ; i < len(strs) ; i++ {
        s := strs[i]
        newMap := make(map[byte]int)
        for j := 0 ; j < len(s) ; j++ {
            newMap[s[j]]++
        }
        jsonMap,_ := json.Marshal(newMap)
        key := string(jsonMap)
        charMap[key] = append(charMap[key],s)
    }

    for _,v := range charMap {
        result = append(result,v)
    }
    
    return result
}