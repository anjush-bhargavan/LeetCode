func leastInterval(tasks []byte, n int) int {
  taskMap := make(map[byte]int)
  max,maxCount := 0,0
    for i := 0 ; i < len(tasks) ; i++ {
        taskMap[tasks[i]]++
        if max < taskMap[tasks[i]] {
            max = taskMap[tasks[i]]
            maxCount = 0
        }
        if max == taskMap[tasks[i]] {
            maxCount++
        }
    }
    num := ((n + 1)*(max - 1) ) + maxCount
    if num > len(tasks) {
        return num
    }
    return len(tasks)
}
