package treeset

import (
	"github.com/alexsteele/go-sets/set"
)

// TreeSet is a red-black tree implementation of SortedSet.
type TreeSet struct {
	root   *node
	cmp    set.Comparator
	length int
}

type colorT bool

const (
	red   = colorT(true)
	black = colorT(false)
)

type node struct {
	color      colorT
	elem       interface{}
	parent     *node
	leftChild  *node
	rightChild *node
}

var nilNode = &node{color: black}

func New(cmp set.Comparator) *TreeSet {
	return &TreeSet{
		root: nilNode,
		cmp:  cmp,
	}
}

func (t *TreeSet) Add(elem interface{}) interface{} {
	curr, parent := t.root, t.root
	var cmp int
	for curr != nilNode {
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

	toAdd := &node{
		color:      red,
		elem:       elem,
		parent:     parent,
		leftChild:  nilNode,
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
	t.rbInsertFixup(toAdd)
	return nil
}

func (t *TreeSet) rbInsertFixup(node *node) {
	if node.parent == nilNode {
		node.color = black
		t.root = node
	} else if node.parent.color == black {
		// Tree is valid.
	} else if uncle := getUncle(node); uncle.color == red {
		node.parent.color = black
		uncle.color = black
		node.parent.parent.color = red
		t.rbInsertFixup(node.parent.parent)
	} else {
		if node.parent == node.parent.parent.leftChild && node == node.parent.rightChild {
			rotateLeft(node.parent)
			node = node.leftChild
		} else if node.parent == node.parent.parent.rightChild && node == node.parent.leftChild {
			rotateRight(node.parent)
			node = node.rightChild
		}

		node.parent.color = black
		node.parent.parent.color = red
		if node == node.parent.leftChild {
			rotateRight(node.parent.parent)
		} else {
			rotateLeft(node.parent.parent)
		}
	}
}

func getUncle(node *node) *node {
	grandparent := node.parent.parent
	if node.parent == grandparent.leftChild {
		return grandparent.rightChild
	} else {
		return grandparent.leftChild
	}
}

func rotateLeft(node *node) {
	node.parent.leftChild = node.rightChild
	node.rightChild.parent = node.parent
	node.parent = node.rightChild
	node.rightChild = node.rightChild.leftChild
	if node.rightChild != nilNode {
		node.rightChild.parent = node		
	}
	node.parent.leftChild = node
}

func rotateRight(node *node) {
	node.parent.rightChild = node.leftChild
	node.leftChild.parent = node.parent
	node.parent = node.leftChild
	node.leftChild = node.leftChild.rightChild
	if node.leftChild != nilNode {
		node.leftChild.parent = node		
	}
	node.parent.rightChild = node
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

func (t *TreeSet) Union(other set.Set) set.Set {
	return nil
}

func (t *TreeSet) Intersect(other set.Set) set.Set {
	return nil
}

func (t *TreeSet) Iter() set.Iterator {
	return nil
}

func (t *TreeSet) ToSlice() []interface{} {
	return nil
}

func (t *TreeSet) String() string {
	return "SortedSet"
}

type iterator struct{}
func (i *iterator) Next() (interface{}, bool) {
	return nil, false
}
