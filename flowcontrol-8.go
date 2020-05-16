package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var before, after float64
	for i := 1; i <= 10; i++ {
		before = z
		fmt.Printf("%d => %v\n", i, before)
		z -= (z*z - x) / (2 * z)
		after = z
		fmt.Printf("%d => %v\n", i, after)

		var sabun float64
		if after > before {
			sabun = after - before
		} else {
			sabun = before - after
		}

		if sabun < 0.00001 {
			fmt.Printf("%d loop break : %f\n", i, sabun)
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
