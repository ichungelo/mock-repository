package mock

import "fmt"

func Square(t int) {
	for i := 0; i < t; i++ {
		for j := 0; j < t; j++ {
			fmt.Print("# ")
		}
		fmt.Println()
	}
	fmt.Println()
}
