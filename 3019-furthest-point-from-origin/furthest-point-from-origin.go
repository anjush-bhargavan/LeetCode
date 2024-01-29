func furthestDistanceFromOrigin(moves string) int {
    distL,distR := 0,0

    for i := 0 ; i < len(moves) ; i++ {
        if moves[i] == 'L' {
            distL++
            distR--
        }else if moves[i] == 'R' {
            distR++
            distL--
        }else{
            distL++
            distR++
        }
    }
    if distL > distR {
        return distL
    }
    return distR
}