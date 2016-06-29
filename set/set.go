package set

// Set represents a mathematical set, a collection which tracks the absence or
// presence of values. Its three core operations, Add, Remove, and Contains, are
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
	String() string 
}

type Iterator interface {
	Next() (interface{}, bool)
}

// SortedSet represents a set whose elements fall under a total order -
// each has a place relative to all others in the set.
type SortedSet interface {
	Set
	
	First() (interface{}, bool)
	Last() (interface{}, bool)
}

type Comparator func(a interface{}, b interface{}) int
