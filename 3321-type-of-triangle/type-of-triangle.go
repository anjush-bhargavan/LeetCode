func triangleType(nums []int) string {
    sort.Ints(nums)
    if nums[0] == nums[1] && nums[1] == nums[2] {
        return "equilateral"
    }
    if (nums[0] == nums[1] && 2*nums[1] > nums[2] ) || nums[1] == nums[2] {
        return "isosceles"
    }
    if (nums[0] + nums[1] > nums[2]) {
        return "scalene"
    }
    return "none"
}