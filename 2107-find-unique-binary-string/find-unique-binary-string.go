func findDifferentBinaryString(nums []string) string {
    binMap := make(map[int]bool)

    for _, s := range nums {
        i, _ := strconv.ParseInt(s, 2, 64)
        in := int(i)
        binMap[in] = true
    }
    for i := 0 ; i < int(math.Pow(float64(2),float64(len(nums)))) ; i++ {
        if !binMap[i] {
            in := int64(i)
            s:= strconv.FormatInt(in, 2)
            for len(s) < len(nums) {
                s = "0" + s
            }
            return s
        }
    }
    return ""
}