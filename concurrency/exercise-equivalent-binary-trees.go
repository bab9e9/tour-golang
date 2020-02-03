package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var n *tree.Tree
	fmt.Println("<Walk(t,ch)>")
	if n = t.Left; n != nil {
		fmt.Print("L")
		Walk(n, ch)
	}
	fmt.Print("V(")
	v := t.Value
	fmt.Printf("%v", v)
	ch <- t.Value
	fmt.Print(")")

	if n = t.Right; n != nil {
		fmt.Print("R")
		Walk(n, ch)
	}
	fmt.Println("</Walk(t,ch)>")
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walker(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walker(t.Left, ch) // simpler, but causes "unecessary" call to Walker(nil, ch) sometimes.
	ch <- t.Value
	Walker(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	/* 2. Test the Walk function.
	        // The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree
			// holding the values k, 2k, 3k, ..., 10k.
			// Create a new channel ch and kick off the walker:
			//	go Walk(tree.New(1), ch)
			// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
	*/

	ch := make(chan int, 100)
	fmt.Println("<go Walk(tree.New(2), ch)>")
	go Walk(tree.New(2), ch)
	fmt.Println("</go Walk(tree.New(2), ch)>")

	ch3 := make(chan int, 100)
	fmt.Println("<go Walker(tree.New(2), ch)>")
	go Walker(tree.New(3), ch3)
	fmt.Println("</go Walker(tree.New(2), ch)>")

	fmt.Println("Check ch")
	for v := range ch {
		fmt.Printf("ch#%d, ", v)
	}

	fmt.Println("Same(tree.New(1), tree.New(1)", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2))", Same(tree.New(1), tree.New(2)))
}

