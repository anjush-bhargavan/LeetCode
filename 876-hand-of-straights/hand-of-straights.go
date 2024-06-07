func isNStraightHand(hand []int, groupSize int) bool {
       if len(hand) % groupSize != 0 {
        return false
    }
    cardMap := make(map[int]int)
    sort.Ints(hand)

    for _,card := range hand {
        cardMap[card]++
    }
    
    for  len(hand) > 0 {
        if cardMap[hand[0]] < 1 {
            hand = hand[1:]
        }else{
            count,n := 0,hand[0]
            for count < groupSize {
                if cardMap[n] > 0 {
                    cardMap[n]--
                    n++
                    count++
                }else{
                    return false
                }
            }
        }

    }

    return len(hand) == 0
}