func hammingWeight(num uint32) int {
    bit := strconv.FormatInt(int64(num),2)
    onebit := strings.Replace(bit,"0","",-1)
    return len(onebit)
}