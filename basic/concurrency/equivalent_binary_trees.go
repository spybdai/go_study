
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Resolve() {
	t1 := tree.New(1)
	t2 := tree.New(-2)
	
	fmt.Println(t1.String())
	fmt.Println(t2.String())
	
	fmt.Println(Same(t1, t2))	
	fmt.Println(Same(tree.New(4), tree.New(4)))	
}

func Walk(t *tree.Tree, c chan int) {
	if t == nil {
		return	
	}

	if t.Left != nil {
		Walk(t.Left, c)
	}

	c <- t.Value

	if t.Right != nil {
		Walk(t.Right, c)
	}
	
}

func Same(t1, t2 *tree.Tree) bool {
	
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	
	go Walk(t1, c1)
	go Walk(t2, c2)

	for count:=0; count<10; count++ {

		v1 := <- c1
		v2 := <- c2 

		if v1 != v2 {
			return false
		}
		
	}
	return true
}
