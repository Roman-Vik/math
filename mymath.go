package mymath

import "math"

func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}
func Abs(x float64)float64{
	return math.Abs(x)
}
func Max(x ...float64) float64 {
	if len(x) == 0 {
		return 0
	}
	max := x[0]
	for _, val := range x[1:] {
		max = math.Max(max, val)
	}
	return max
}
 func Yn(n int, x float64) float64{
	return math.Yn(n,x)
}
