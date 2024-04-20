// В файле mymath.go
package mymath

import "math"

// Sqrt возвращает квадратный корень из числа x.
func Sqrt(x float64) int {
    return int(math.Sqrt(x))
}
