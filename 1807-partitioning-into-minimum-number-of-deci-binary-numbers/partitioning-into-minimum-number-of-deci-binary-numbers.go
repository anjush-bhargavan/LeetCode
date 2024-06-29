func minPartitions(n string) int {
    max := 0
    for _,s := range n {
        num := int(s - '0')
        if max < num {
            max = num
        }
    }
    return max
}