func tribonacci(n int) int {
    x,y,z := 0,1,1
    for i := 0 ; i < n ; i++ {
        x,y,z = y,z,x+y+z
    }
    return x
}