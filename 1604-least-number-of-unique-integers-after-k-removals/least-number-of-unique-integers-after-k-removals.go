func findLeastNumOfUniqueInts(arr []int, k int) int {
    intMap := make(map[int]int)

    for i := 0 ; i < len(arr) ; i++ {
        intMap[arr[i]]++
    }

    sortedArr := sortValue(intMap)

    count := 0
    for i := 0 ; k > 0 ; i++ {
       if k >= intMap[sortedArr[i]] {
           k -= intMap[sortedArr[i]]
           count++
       }else{
           break
       }
    }
  
    return len(intMap) - count

}


func sortValue (mp map[int]int) []int{
    var sortedKeys []int
    for keys:= range mp{
        sortedKeys=append(sortedKeys,keys)
    }
    sort.Slice(sortedKeys,func(i,j int) bool{
        return mp[sortedKeys[i]]<mp[sortedKeys[j]]
    })
    return sortedKeys
}