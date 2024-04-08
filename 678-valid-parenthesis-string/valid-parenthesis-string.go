func checkValidString(s string) bool {
    countMin,countMax := 0,0
    for _,letter := range s {
        if letter == '(' {
            countMin++
            countMax++
        }else if letter == ')' {
            countMin--
            countMax--
        }else if letter == '*' {
            countMin--
            countMax++
        }
        if countMax < 0 {
            return false
        }
        if countMin  < 0 {
            countMin = 0
        }
    }
    return countMin == 0 
}