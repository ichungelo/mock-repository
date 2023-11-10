package mock

func Coin(n int, k int) int {
	if k == 0 {
		return 1
	}

	if k < 0 {
		return 0
	}

	if n <= 0 {
		return 0
	}

	return Coin(n-1, k) + Coin(n, k-n)
}