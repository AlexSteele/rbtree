package rbtree

import (
	"fmt"
	"math"
)

func (t *RBTree) printLevelOrder() {
	q := []*node{t.root}
	count := 0
	for len(q) > 0 {
		
		curr := q[0]
		q = q[1:]
		count++
		
		if level := math.Log2(float64(count)); math.Ceil(level) == level {
			fmt.Println("\n---------------Level---------------: ", level, "\n")
		}
		
		if curr != nilNode {
			fmt.Println("\t\tNode: ", curr)
			q = append(q, curr.leftChild, curr.rightChild)
		}
	}
}
