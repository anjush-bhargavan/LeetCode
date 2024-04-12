func trap(height []int) int {
    maxRight := make([]int,len(height))
    max,count := 0,0
    for i := len(height)-2 ; i >= 0 ; i-- {
        if max < height[i+1] {
            max = height[i+1]
        }
        maxRight[i] = max
    }

    max = 0
    for i,h := range height {
        water := min(max,maxRight[i]) - h
        if water > 0 {
            count += water
        }
        if max < h {
            max = h
        }
    }
    
    return count
}