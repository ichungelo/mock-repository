package mock

import "fmt"

func Triangle(l int)  {
	for i := 0; i < l; i++ {
		for j := 0; j < l-i; j++ {
			fmt.Print("# ")
		}
		fmt.Println()
	}
	fmt.Println()
}