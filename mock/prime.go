package mock

import "math"

func Prime(total int) []int {
	result := []int{}
	i := 2
loop1:
	for len(result) < total {
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				i++
				continue loop1
			}
		}
		result = append(result, i)
		i++
	}
	return result
}
