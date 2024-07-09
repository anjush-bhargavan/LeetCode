func averageWaitingTime(customers [][]int) float64 {
    time,total := 0,0

    for i := 0 ; i < len(customers) ; i++ {
        if time < customers[i][0] {
            time = customers[i][0]
        }
        time += customers[i][1]
        total += (time - customers[i][0])
    }
    return float64(total)/float64(len(customers))
}