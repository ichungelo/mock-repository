package mock

func IntReverser(num int) int {
	res := 0
	for num > 0 {
		remain := num % 10
		res = (res * 10) + remain
		num/=10
	}
	return res
}
