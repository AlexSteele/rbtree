package set

// TODO: Provide Union, Intersection operations.

// Set represents an unordered collection of values where duplicates are disallowed.
type Set interface {
	Add(elem interface{}) interface{} 
	Remove(elem interface{}) bool 
	Contains(elem interface{}) bool
	Length() int
	IsEmpty() bool 
	Clear() 
	Iter() Iterator
	ToSlice() []interface{} 
	String() string 
}

type Iterator interface {
	Next() (interface{}, bool)
}

// SortedSet represents a totally ordered set. In addition to implementing the Set interface
// it guarantees that that Iter and ToSlice will return values in sorted order
// and provides First and Last, which return the first and last elements under the ordering.
type SortedSet interface {
	Set
	
	First() (interface{}, bool)
	Last() (interface{}, bool)
}
