package set

import (
	"fmt"
)

// TODO: Read up on interface equality.

// Set represents a mathematical set, a collection which tracks the absence
// or presence of values. Its three core operations, Add, Remove, and Contains, are
// guaranteed to run in constant time.
type Set interface {
	Add(elem interface{}) interface{} 
	Remove(elem interface{}) bool 
	Contains(elem interface{}) bool
	Clear() 
	Length() int
	IsEmpty() bool 
	Union(other Set) Set
	Intersect(other Set) Set
	ToSlice() []interface{} 
	Iter() Iterator
	fmt.Stringer // TODO: This may not belong
}

type Iterator interface {
	Next() (interface{}, bool)
}

type iteratorImpl struct {
	elems  []interface{}
	length int
	pos    int 
}

func (i *iteratorImpl) Next() interface{}, bool {
	if i.pos < i.length {
		elem := i.elems[pos]
		i.pos += 1
		return elem, true
	}
	return nil, false
}

type HashSet struct {
	elems map[interface{}]struct{} 
}

func New() *HashSet {
	return &HashSet{
		elems: make(map[interface{}]struct{}, 0) 
	}
}

func (h *HashSet) Add(elem interface{}) interface{} {
	old, _ := h.elems[elem]
	h.elems[elem] = struct{}{}
	return old 
}

func (h *HashSet) Remove(elem interface{}) bool {
	_, exists := h.elems[elem]
	if exists {
		 delete(h.elems, elem) 
	}
	return exists 
}

func (h *HashSet) Contains(elem interface{}) bool {
	_, exists := h.elems[elem]
	return exists
}

func (h *HashSet) Clear() {
	h.elems = make(map[interface{}]struct{}, 0) 
}

func (h *HashSet) Length() int {
	return len(h.elems) 
}

func (h *HashSet) IsEmpty() bool {
	return h.Size() == 0 
}

func (h *HashSet) Union(other Set) Set {
	new := make(map[interface{}]struct{}, len(h.elems) + len(other.Size())) // TODO: Sum of the lengths is likely inappropriate.

	// Note that this strategy gives preference to the elems in the second set. 
	for elem, _ := range h.elems {
		new[elem] = struct{}{} 
	}
	iter := other.Iter()
	for elem, exists := iter.Next(); exists; elem, exists := iter.Next() {
		new[elem] = struct{}{} 
	}

	return &HashSet{new}
}

func (h *HashSet) Intersect(other Set) Set {
	new := make(map[interface{}]struct{}, len(h.elems))

	for elem, _ := range h.elems {
		if other.Contains(elem) {
			new[elem] = struct{}{} 
		}
	}

	return &HashSet{new}
}

func (h *HashSet) ToSlice() []interface{} {
	slice := make([]interface{}, len(h.elems))

	for elem, _ := range h.elems {
		slice = append(slice, elem) 
	}

	return slice
}

func (h *HashSet) Iter() Iterator {
	elems := h.ToSlice() 
	return &iteratorImpl{
		elems: elems,
		length: len(elems),
		pos: 0,
	}
}

func (h *HashSet) String() string {
	return "HashSet"
}
