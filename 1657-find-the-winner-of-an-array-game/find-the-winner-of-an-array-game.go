func getWinner(arr []int, k int) int {
    count,winner := 0,arr[0]
    if k >= len(arr) {
        sort.Ints(arr)
        return arr[len(arr)-1]
    } 
    for i := 0 ; i < len(arr) ; i++ {
        if arr[0] > arr[1] {
            arr[0],arr[1] = arr[1],arr[0]
        }
            arr = append(arr,arr[0])
            arr = arr[1:]
            count++
            if winner != arr[0] {
                count = 1 
                winner = arr[0]
            }
            if count == k {
                return winner
            }
    }
    return arr[0]
}