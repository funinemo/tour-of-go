package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := -1
	var p1, p2 int
	return func() int {
		switch fib {
		case -1:
			p2 = 1
			p1 = 0
			fib = 0
			return 0
		}
		fib = p1 + p2
		p2 = p1
		p1 = fib
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
