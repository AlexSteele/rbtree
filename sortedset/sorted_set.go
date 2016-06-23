package sortedset

import (
	"github.com/alexsteele/go-sets/set"
)

// SortedSet represents a set whose elements fall under a total order -
// each has a place relative to all others in the set. 
type SortedSet interface {
	set.Set
	First() (interface{}, bool)
	Nth(i int) (interface{}, bool)
	Last() (interface{}, bool)
}

type Comparator func(interface{}, interface{}) int 

type sortedSetIter struct {}

func (i *sortedSetIter) Next() (interface{}, bool) {
	return nil, false
}

// TreeSet is a red-black tree implementation of SortedSet. 
type TreeSet struct {
	root *treeNode
	cmp Comparator
	length int 
}

type color int

const (
	red color = iota
	black 
)

type treeNode struct {
	c color
	elem interface{}
	parent *treeNode 
	leftChild *treeNode
	rightChild *treeNode 
}

const nilNode = &treeNode{
	c: black,
}

func New(cmp Comparator) SortedSet {
	return &TreeSet{
		root: nilNode,
		cmp: Comparator,
	} 
}

func (t *TreeSet) Add(elem interface{}) interface{} {
	curr, parent := t.root, t.root
	var cmp int
	for curr := nil {
		parent = curr 
		cmp = t.cmp(elem, curr)
		if cmp == 0 {
			old := curr.elem 
			curr.elem = elem
			return old 
		} else if cmp < 0 {
			curr = curr.leftChild
		} else {
			curr = curr.rightChild
		}
	}

	toAdd := &treeNode{
		c: red,
		elem: elem,
		parent: parent,
		leftChild: nilNode,
		rightChild: nilNode, 
	}

	if parent != nilNode {
		if cmp < 0 {
			parent.leftChild = toAdd
		} else {
			parent.rightChild = toAdd
		}		
	}

	t.length += 1
	rbInsertFixup(toAdd)
	return nil
}

func rbInsertFixup(node *treeNode) {
	parent := node.parent 
	if parent == nilNode {
		t.root = node
	} else if parent.c == black {
		// Tree is valid.
	} else if uncle := getUncle(node); uncle.c == red {
		parent.c = black
		uncle.c = black
		parent.parent = red
		rbInsertFixup(parent.parent)
	} else {
		grandparent := node.parent.parent
		if (parent == grandparent.leftChild && node == parent.rightChild) {
			rotateLeft(parent)
			node = node.leftChild
		} else if (parent == grandparent.rightChild && node == parent.leftChild) {
			rotateRight(parent)
			node = node.rightChild
		}

		node.parent.c = black
		node.parent.parent.c = red
		if node == node.parent.leftChild {
			rotateRight(parent.parent)
		} else {
			rotateLeft(parent.parent) 
		}
	}	
}

// Assumes that the grandparent is not nilNode. Will return nilNode if the uncle
// is nil. 
func getUncle(node *treeNode) *treeNode {
	grandparent := node.parent.parent
	if node.parent == grandparent.leftChild {
		return grandparent.rightChild 
	} else {
		return grandparent.leftChild 
	}
}

// Assumes that the grandparent is not nilNode and that node.leftChild is the
// nilNode. 
func rotateLeft(node *treeNode) {
	node.leftChild = node.parent
	if node.parent == node.parent.parent.leftChild {
		node.parent.parent.leftChild = node
	} else {
		node.parent.parent.rightChild = node
	}
	node.parent = node.parent.parent
	node.leftChild.parent = node 
}

// Assumes that the grandparent is not the nilNode and that node.rightChild is
// the nilNode.
func rotateRight(node *treeNode) {
	node.rightChild = node.parent
	if node.parent == node.parent.parent.leftChild {
		node.parent.parent.leftChild = node
	} else {
		node.parent.parent.rightChild = node
	}
	node.parent = node.parent.parent
	node.rightChild.parent = node 
}

func (t *TreeSet) Remove(elem interface{}) bool {
	return false
}

func (t *TreeSet) Contains(elem interface{}) bool {
	curr := t.root
	for curr != nilNode {
		cmp := t.cmp(elem, curr) 
		if cmp == 0 {
			return true 
		} else if cmp < 0 {
			curr = curr.leftChild
		} else {
			curr = curr.leftChild
		}
	}
	return false 
}

func (t *TreeSet) Clear() {
	t.root = nilNode
	t.length = 0
}

func (t *TreeSet) Length() int {
	return t.length
}

func (t *TreeSet) IsEmpty() bool {
	return t.length == 0 
}

func (t *TreeSet) Union(other Set) Set {
	return nil 
}

func (t *TreeSet) Intersect(other Set) Set {
	return nil
}

func (t *TreeSet) Iter() Iterator {
	return nil 
}

func (t *TreeSet) ToSlice() []interface{} {
	return nil
}

func (t *TreeSet) String() string {
	return "SortedSet" 
}

func (t *TreeSet) First() (interface{}, bool) {
	if t.root == nilNode {
		return nil, false 
	}
	curr := t.root
	for curr.leftChild != nilNode {
		curr = curr.leftChild 
	}
	return curr, true 
}

func (t *TreeSet) Nth(i int) (interface{}, bool) {
	return nil, false 
}

func (t *TreeSet) Last() (interface{}, bool) {
	if t.root == nilNode {
		return nil, false
	}
	curr := t.root
	for curr.rightChild != nilNode {
		curr = curr.rightChild
	}
	return curr, true 
}
