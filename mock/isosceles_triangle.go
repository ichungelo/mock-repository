package mock

import "fmt"

// SEGITIGA SAMA KAKI
func IsoscelesTriangle(l int) {
	fmt.Println("########## START ISOSCELES TRIANGLE ##########")
	for i := 1; i <= l; i++ {
		for j := 0; j < l-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*i-1; k++ {
			fmt.Print("#")
		}

		fmt.Println()
	}
}
