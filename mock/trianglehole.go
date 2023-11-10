package mock

import "fmt"

func TriangleHole(l int)  {
	for i := 0; i < l; i++ {
		if i > 0 && i < l - 1 {
			fmt.Print("# ")
			for k := 0; k < l-i-2; k++ {
				fmt.Print("  ")
			}
			fmt.Print("# ")
		} else {
			for j := 0; j < l-i; j++ {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}