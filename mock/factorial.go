package mock

func Factorial(input uint64) uint64 {
	if input == 1 || input == 0 {
		return input
	}
	return input * Factorial(input-1)
}
