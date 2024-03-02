func complexNumberMultiply(num1 string, num2 string) string {
    c1,_ := strconv.ParseComplex(num1,64)
    c2,_ := strconv.ParseComplex(num2,64)
    c3 := c1 * c2
    real := int(real(c3))
    img := int(imag(c3))
    
    return fmt.Sprintf("%d+%di",real,img)
}