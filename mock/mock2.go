package mock

import "strconv"

func FizzBuzz(max int) []string {
	res := []string{}
	for i := 1; i <= max; i++ {
		if i%5 == 0 || i%5 == 2 {
			res = append(res, "*")
		} else {
			strNum := strconv.Itoa(i)
			res = append(res, strNum)
		}
	}

	return res
}
