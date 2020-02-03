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
	ch1 := make(chan int, 100)
	go Walk(t1, ch1)
	close(ch1) // or range will block

	ch2 := make(chan int, 100)
	go Walk(t2, ch2)
	close(ch2) // or range will block

	return SameCh(ch1, ch2)
}

// Same determines whether t1 and t2 contain the same values.
func SameCh(ch1, ch2 chan int) bool {
	for {
		v1, b1 := <-ch1
		v2, b2 := <-ch1

		if !(b1 || b2) { // empty
			fmt.Printf("!(b1 || b2): %v, %v, %v, %v\n", b1, b2, v1, v2)
			return true
		}

		if b1 != b2 { // different lengths
			fmt.Printf("b1 != b2:  %v, %v, %v, %v\n", b1, b2, v1, v2)
			return false
		}

		if v1 != v2 { // different values
			fmt.Printf("v1 != v2:  %v, %v, %v, %v\n", b1, b2, v1, v2)
			return false
		}
	}
	fmt.Println("unreachable?")
	return false // unreachable
}

func TestWalk(walk func(t *tree.Tree, ch chan int)) {
	/* 2. Test the Walk function.
	        // The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree
			// holding the values k, 2k, 3k, ..., 10k.
			// Create a new channel ch and kick off the walker:
			//	go Walk(tree.New(1), ch)
			// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
	*/

	ch2 := make(chan int, 100)
	fmt.Println("<go walk(tree.New(2), ch2)>")
	go walk(tree.New(2), ch2)
	fmt.Println("</go walk(tree.New(2), ch2)>")

	fmt.Println("Check ch2")
	close(ch2) // or range will block
	for v := range ch2 {
		fmt.Printf("ch2#%d, ", v)
	}
	return
}

func main() {
	TestWalk(Walk)
	TestWalk(Walker)

	// Test Same
	fmt.Println("Same(tree.New(1), tree.New(1)", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2))", Same(tree.New(1), tree.New(2)))
}
