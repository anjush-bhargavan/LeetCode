func wateringPlants(plants []int, capacity int) int {
    steps,can := 0,capacity
    for i := 0 ; i  < len(plants) ; i++ {
        if plants[i] <= can {
            can -= plants[i]
            steps++
        }else{
            can = capacity - plants[i]
            steps += (2*i)+1   
        }
    }
    return steps
}