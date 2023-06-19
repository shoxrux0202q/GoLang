package homework
import "fmt"

func IsDiv8(a int) int {
    var b, sum int
    for i := 1; i <= a; i++ {
        fmt.Scan(&b)
        if 10 <= b && b%8 == 0 && b <= 99 {
            sum += b
        }
    }
    return sum
}

