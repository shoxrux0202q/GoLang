package homework
func Double_m(a, b int) int {
    sum := 0
    for i := a; i <= b; i++ {
        square := i * i
        sum = sum + square
    }
    return sum
}