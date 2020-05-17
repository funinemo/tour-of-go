package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// 参照 https://qiita.com/rock619/items/f412d1f870a022c142d0

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	//	fmt.Printf("---%d : %v\n",t.Value,t)
	walk(t, ch)
	close(ch)

}
func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		d1, ok1 := <-ch1
		d2, ok2 := <-ch2
		if ok1 == true || ok2 == true {
			if d1 != d2 {
				return false
			}
		}
		break
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Printf("ch read : %d\n", i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
