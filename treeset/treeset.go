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
	red colorT   = true
	black colorT = false
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
		cmp = t.cmp(elem, curr.elem)
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

	t.rbInsertFixup(toAdd)
	t.length += 1
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
	if node == node.parent.leftChild {
		node.parent.leftChild = node.rightChild				
	} else {
		node.parent.rightChild = node.rightChild
	}
	node.rightChild.parent = node.parent
	node.parent = node.rightChild
	node.rightChild = node.rightChild.leftChild
	if node.rightChild != nilNode {
		node.rightChild.parent = node		
	}
	node.parent.leftChild = node
}

func rotateRight(node *node) {
	if node == node.parent.leftChild {
		node.parent.leftChild = node.leftChild
	} else {
		node.parent.rightChild = node.leftChild
	}
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
	toRemove := t.getNode(elem)
	if toRemove == nil {
		return false
	}
	
	if successor := getSuccessor(toRemove); successor != nilNode {
		toRemove.elem = successor.elem
		toRemove = successor
	}
	
	// toRemove has either 1 or 0 non-nil children. Replace
	// toRemove with its child.
	var child *node
	if toRemove.leftChild == nilNode {
		child = toRemove.rightChild
	} else {
		child = toRemove.leftChild
	}
	if toRemove == toRemove.parent.leftChild {
		toRemove.parent.leftChild = child
	} else {
		toRemove.parent.rightChild = child
	}
	if child != nilNode {
		child.parent = toRemove.parent
	}

	// Restore the tree's invariants.
	if toRemove.color == red {
		// toRemove is not the root. We're done.
	} else if child.color == red {
		// toRemove is not the root. It is black and child is red. 
		child.color = black
	} else {
		t.rbRemoveFixup(child)
	}
	
	t.length -= 1
	return true
}

func (t *TreeSet) rbRemoveFixup(child *node) {
	
	// Both toRemove and child are black (with child possibly nilNode)
	// toRemove could be the root.
	// toRemove could be the successor to the thing we actually want to remove.

	// TODO: WHAT IF CHILD IS NILNODE?
	if child == nilNode {
		return
	}

	for {
		// TODO: MAY BE if "child.parent == nil"
		if child.parent == nilNode {
			// child is the new root.
			t.root = child
			return 
		}
		
		var sibling *node
		if child == child.parent.leftChild {
			sibling = child.parent.rightChild
		} else {
			sibling = child.parent.leftChild
		}

		if sibling.color == red {
			child.parent.color = red
			sibling.color = black
			if child == child.parent.leftChild {
				rotateLeft(child.parent)
				sibling = child.parent.rightChild // TODO: HOT-SPOT. Sibling may be incorrect (here and below).
			} else {
				rotateRight(child.parent)
				sibling = child.parent.leftChild
			}
		}
		if sibling.color == black &&
			sibling.leftChild.color == black &&
			sibling.rightChild.color == black {

			sibling.color = red
			if child.parent.color == black {
				
				// Repeat the fixup with the parent.
				child = child.parent
				continue 
			} else {
				child.parent.color = black
			}
		} else {
			if sibling.color == black {
				if child == child.parent.leftChild &&
					sibling.rightChild.color == black &&
					sibling.leftChild.color == red {

					rotateRight(sibling)
				} else if child == child.parent.rightChild &&
					sibling.leftChild.color == black &&
					sibling.rightChild.color == red {

					sibling.color = red
					sibling.rightChild.color = black
					rotateLeft(sibling)
				}
			}

			sibling.color = child.parent.color
			child.parent.color = black

			if child == child.parent.leftChild {
				sibling.rightChild.color = black
				rotateLeft(child.parent)
			} else {
				sibling.leftChild.color = black
				rotateRight(child.parent)
			}
		}
		
		return
	}
}

func (t *TreeSet) Contains(elem interface{}) bool {
	return t.getNode(elem) != nil
}

// Returns nil if no node with the given element exists.
func (t *TreeSet) getNode(elem interface{}) *node {
	curr := t.root
	for curr != nilNode {
		cmp := t.cmp(elem, curr.elem)
		if cmp == 0 {
			return curr
		} else if cmp < 0 {
			curr = curr.leftChild
		} else {
			curr = curr.leftChild
		}
	}
	return nil
}

func getSuccessor(n *node) *node {
	curr := n.rightChild
	if curr == nilNode {
		return curr
	}
	for curr.leftChild != nil {
		curr = curr.leftChild
	}
	return curr 
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
