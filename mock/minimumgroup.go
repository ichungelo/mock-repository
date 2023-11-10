package mock

func MinimumGroup(predators []int) int {
	var (
		max    = 1
		arrLen = len(predators)
	)

	for i := 0; i < len(predators); i++ {
		var (
			counter = 1
			a       = i
		)

		for predators[a] > -1 && predators[a] < arrLen && counter < arrLen {
			a = predators[a]
			counter++
		}

		if counter > max {
			max = counter
		}
	}

	return max
}
