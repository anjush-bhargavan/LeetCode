func findGCD(nums []int) int {
    sort.Ints(nums)
    min,max := nums[0],nums[len(nums)-1]
    for i := min ; i >= 1 ; i-- {
        if Divisor(min,max,i){
            return i
        } 
    }
    return 1
}

func Divisor(a,b,num int) bool {
    return a%num == 0 && b%num == 0
}