func sumIndicesWithKSetBits(nums []int, k int) int {
    sum:=0
    for i:=0;i<len(nums);i++{
        b:=strconv.FormatInt(int64(i),2)
        bNum,_ := strconv.Atoi(b)
        sumB:=0
        for bNum>0{
            sumB+=bNum%10
            bNum/=10
        }
        if sumB==k{
            sum+=nums[i]
        }
    }
    return sum
}