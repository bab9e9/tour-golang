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
	/* 2. Test the Walk function. 
	 The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)

Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
	*/
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Same(tree.New(1), tree.New(1)",Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2))",Same(tree.New(1), tree.New(2)))
}

