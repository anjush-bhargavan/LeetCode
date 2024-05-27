func asteroidCollision(asteroids []int) []int {
	stack := []int{asteroids[0]}

	for i := 1; i < len(asteroids); i++ {
		if len(stack) == 0 || asteroids[i] < 0 && stack[len(stack)-1] < 0 || asteroids[i] > 0 && stack[len(stack)-1] > 0 {
			stack = append(stack, asteroids[i])
		} else {
			for i < len(asteroids) && len(stack) > 0 && -1*asteroids[i] == stack[len(stack)-1] && asteroids[i] < 0 {
				stack = stack[:len(stack)-1]
				i++
			}
			for i < len(asteroids) && len(stack) > 0 && asteroids[i] < 0 && stack[len(stack)-1] > 0 {
				if (-1 * asteroids[i]) > stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else if (-1 * asteroids[i]) == stack[len(stack)-1]{
                    stack = stack[:len(stack)-1]
					i++
				}else{
                    i++
                }
			}
			if i < len(asteroids) {
				stack = append(stack, asteroids[i])
			}
		}
	}
	return stack
}
