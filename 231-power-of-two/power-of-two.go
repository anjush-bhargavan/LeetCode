func isPowerOfTwo(n int) bool {
   return n > 0 && ((1 << 30) % n) == 0
}