package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch
// There's no way that I could have resolved this alone.
func Walk(t *tree.Tree, ch chan int) {
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		if t.Left != nil {
			walk(t.Left)
		}
		ch <- t.Value
		if t.Right != nil {
			walk(t.Right)
		}
	}

	walk(t)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	t1_ch := make(chan int)
	t2_ch := make(chan int)

	go Walk(t1, t1_ch)
	go Walk(t2, t2_ch)

	for i := 0; i < 10; i++ {
		if <-t1_ch != <-t2_ch {
			return false
		}
	}

	return true
}

func main() {
	tree_ch := make(chan int)
	go Walk(tree.New(1), tree_ch)
	for v := range tree_ch {
		fmt.Println(v)
	}

	vrai := Same(tree.New(1), tree.New(1))
	fmt.Println("Same values:", vrai) // true

	faux := Same(tree.New(1), tree.New(2))
	fmt.Println("Different values:", faux) // false
}

// https://go.dev/tour/concurrency/7
// https://go.dev/tour/concurrency/8
