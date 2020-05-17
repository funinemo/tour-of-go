package main

import (
	"fmt"
)

// ErrNegativeSqrt 型
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) float64 {
	if x < 0 {
		fmt.Printf("%s\n", ErrNegativeSqrt(x).Error())
		return x
	}
	z := 1.0
	var before, after float64
	for i := 1; i <= 10; i++ {
		before = z
		z -= (z*z - x) / (2 * z)
		after = z

		var sabun float64
		if after > before {
			sabun = after - before
		} else {
			sabun = before - after
		}

		if sabun < 0.00001 {
			return z
		}
	}
	return z
}

//func Sqrt(x float64) (float64, error) {
//	return 0, nil
//}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
