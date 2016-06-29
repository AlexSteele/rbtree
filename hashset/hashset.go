package hashset

import "github.com/alexsteele/go-sets/set"

type HashSet struct {
	elems map[interface{}]struct{} 
}

func New() *HashSet {
	return &HashSet{
		elems: make(map[interface{}]struct{}, 0),
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
	return h.Length() == 0 
}

func (h *HashSet) Union(other set.Set) Set {
	newSet := make(map[interface{}]struct{}, h.Length() + other.Length()) // TODO: Sum of the lengths is likely inappropriate.

	// Note that this strategy gives preference to the elems in the second set. 
	for elem, _ := range h.elems {
		newSet[elem] = struct{}{} 
	}
	iter := other.Iter()
	for elem, exists := iter.Next(); exists; elem, exists = iter.Next() {
		newSet[elem] = struct{}{} 
	}

	return &HashSet{newSet}
}

func (h *HashSet) Intersect(other set.Set) Set {
	newSet := make(map[interface{}]struct{}, len(h.elems))

	for elem, _ := range h.elems {
		if other.Contains(elem) {
			newSet[elem] = struct{}{} 
		}
	}

	return &HashSet{newSet}
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
	return &iterator{
		elems: elems,
		length: len(elems),
		pos: 0,
	}
}

func (h *HashSet) String() string {
	return "HashSet"
}

type iterator struct {
	elems  []interface{}
	length int
	pos    int 
}

func (i *iterator) Next() (interface{}, bool) {
	if i.pos < i.length {
		elem := i.elems[i.pos]
		i.pos += 1
		return elem, true
	}
	return nil, false
}
