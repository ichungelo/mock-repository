package mock

func LongestTime(keyTime [][2]int) string {
	var (
		key      = keyTime[0][0]
		diffTime = keyTime[0][1]
	)

	for i := 1; i < len(keyTime); i++ {
		newDiff := keyTime[i][1] - keyTime[i-1][1]
		if newDiff > diffTime {
			key = keyTime[i][0]
			diffTime = keyTime[i][1]
		}
	}

	charDec := 97 + key
	return string(rune(charDec))
}
