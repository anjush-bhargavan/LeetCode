func canVisitAllRooms(rooms [][]int) bool {
    keyMap := make(map[int]bool)
    keyMap[0] = true
    i,flag := 0,false
    for {
        if keyMap[i] {
            for j := 0 ; j < len(rooms[i]) ; j++ {
                if !keyMap[rooms[i][j]] {
                    keyMap[rooms[i][j]] = true
                    flag = true
                }
            }
        }
        i++
        if i == len(rooms) {
            if flag {
                i = 0
                flag = false
            }else{
                break
            }
        }
    }
    return len(keyMap) == len(rooms)
}