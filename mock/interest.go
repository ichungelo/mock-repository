package mock

import "math"

func InterestPerYear(year int, interest float64, amount float64) float64 {
	return math.Pow(interest/100+1, float64(year)) * amount
}
