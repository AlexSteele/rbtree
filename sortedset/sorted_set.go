package sortedset

import (
	"github.com/alexsteele/go-sets/set"
)

// SortedSet represents a set whose elements fall under a total order -
// each has a place relative to all others in the set. 
type SortedSet interface {
	set.Set
	First() interface{}, bool
	Nth(i int) interface{}, bool
	Last() interface{}, bool
	SetComparator(func(interface{}, interface{}) bool) 
}

type sortedSetIter struct {}

func (i *sortedSetIter) Next() (interface{}, bool) {
	return nil, false
}

// TreeSet is a BST implementation of SortedSet. 
type TreeSet struct {
	
}

func New() *TreeSet {
	return &TreeSet{} 
}

func Add(elem interface{}) interface{} {
	return nil
}

func Remove(elem interface{}) bool {
	return false
}

func Contains(elem interface{}) bool {
	return false
}

func Clear() {
	return 
}

func Length() int {
	return 0
}

func (s *SortedSet) IsEmpty() bool {
	return true 
}

func (s *SortedSet) Union(other Set) Set {
	return nil 
}

func (s *SortedSet) Intersect(other Set) Set {
	return nil
}

func (s *SortedSet) Iter() Iterator {
	return nil 
}

func (s *SortedSet) ToSlice() []interface{} {
	return nil
}

func (s *SortedSet) String() string {
	return "SortedSet" 
}

func (s *SortedSet) First() (interface{}, bool) {
	return nil, false
}

func Nth(i int) (interface{}, bool) {
	return nil, false 
}

func (s *SortedSet) Last() (interface{}, bool) {
	return nil, false 
}
