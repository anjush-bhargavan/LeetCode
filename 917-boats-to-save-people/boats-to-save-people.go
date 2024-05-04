func numRescueBoats(people []int, limit int) int {
    count,i,j := 0,0,len(people)-1
    sort.Ints(people)
    for i <= j {
        if people[i] + people[j] <= limit {
            i++
            j--
        }else{
            j--
        }
        count++
    }
    return count
}