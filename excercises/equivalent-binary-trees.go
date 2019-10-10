package main

import "fmt"
import "golang.org/x/tour/tree"

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkTree(t, ch)
	close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}
	return true
}

func printSame(b bool) {
	if b == true {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are NOT the same")
	}
}

func main() {
	t1 := tree.New(1)
	fmt.Println(t1)
	ch := make(chan int)
	go Walk(t1, ch)
	for i := range ch {
		fmt.Println(i)
	}
	printSame(Same(tree.New(1), tree.New(1)))
	printSame(Same(tree.New(1), tree.New(2)))
}
