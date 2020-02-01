package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false // test should fail
}

func main() {
	fmt.Println("Same(tree.New(1), tree.New(1)",Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2))",Same(tree.New(1), tree.New(2)))
}

