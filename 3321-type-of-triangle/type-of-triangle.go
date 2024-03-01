func triangleType(nums []int) string {
    a,b,c := nums[0],nums[1],nums[2]
    if a + b <= c || b + c <= a || a + c <= b {
        return "none"
    }
    if a == b && b == c {
        return "equilateral"
    }
    if a == b  || b == c || a == c {
        return "isosceles"
    }
    return "scalene"
}