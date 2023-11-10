package mock

func GetNumber(arr []int, target int) [2]int {
	var (
		mapStorage map[int]int = make(map[int]int)
		result     [2]int
	)

	for _, v := range arr {
		value, ok := mapStorage[v]
		if ok {
			result = [2]int{
				value, v,
			}

			break
		}

		mapStorage[target-v] = v
	}

	return result
}
