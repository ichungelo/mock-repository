package mock

import "math"

func SumOfSeries(n int) float64 {
	return math.Sqrt(float64(n+1)) - 1
}
