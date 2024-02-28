func canVisitAllRooms(rooms [][]int) bool {
    keyMap := make(map[int]bool)
    DFS(rooms,keyMap,0)
    return len(keyMap) == len(rooms)
}


func DFS(arr [][]int,keyMap map[int]bool,val int) {
    keyMap[val] = true
    for _,v := range arr[val] {
        if !keyMap[v] {
            DFS(arr,keyMap,v)
        }
    }
}