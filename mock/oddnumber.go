package mock

func OddNumber(total int) []int {
	result := []int{}
	i := 1
	for len(result) < total {
		if i%2 == 1 {
			result = append(result, i)
		}
		i++
	}
	return result
}
