package mock

func Jump (first, second, length int) []int {
	jump := second - first
	res := []int{}

	for i := 0; i < length; i++ {
		res = append(res, first)
		first += jump
	}

	return res
}