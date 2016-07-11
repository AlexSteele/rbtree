package rbtree

import (
	"strings"
)

// Comparator takes two arguments of interface{} type.
// It returns a negative int if the first is less than the second, a positive
// int if it is greater, and 0 if the two are equal.
//
// A panic is expected if arguments of incompatible type are given.
//
// Package rbtree provides implementations for most types defined in package builtin.
type Comparator func(a interface{}, b interface{}) int

var Float32Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(float32) - b.(float32)
	switch {
	case cmp == 0.0:
		return 0
	case cmp > 0.0:
		return 1
	default:
		return -1
	}
}

var Float64Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(float64) - b.(float64)
	switch {
	case cmp == 0.0:
		return 0
	case cmp > 0.0:
		return 1
	default:
		return -1
	}
}

var IntComparator Comparator = func(a interface{}, b interface{}) int {
	return a.(int) - b.(int)
}

var Int16Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(int16) - b.(int16)
	switch {
	case cmp == 0:
		return 0
	case cmp > 0:
		return 1
	default:
		return -1
	}
}

var Int32Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(int32) - b.(int32)
	switch {
	case cmp == 0:
		return 0
	case cmp > 0:
		return 1
	default:
		return -1
	}
}

var Int64Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(int64) - b.(int64)
	switch {
	case cmp == 0:
		return 0
	case cmp > 0:
		return 1
	default:
		return -1
	}
}

var Int8Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(int8) - b.(int8)
	switch {
	case cmp == 0:
		return 0
	case cmp > 0:
		return 1
	default:
		return -1
	}
}

var RuneComparator Comparator = func(a interface{}, b interface{}) int {
	return int(a.(rune) - b.(rune))
}

// StringComparator wraps strings.Compare(a, b), returning the result of a lexicographic
// comparison.
var StringComparator Comparator = func(a interface{}, b interface{}) int {
	return strings.Compare(a.(string), b.(string))
}

var UIntComparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(uint) - b.(uint)
	switch {
	case cmp == 0:
		return 0
	case cmp > 0:
		return 1
	default:
		return -1
	}
}

var UInt16Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(uint16) - b.(uint16)
	switch {
	case cmp == 0:
		return 0
	case cmp > 0:
		return 1
	default:
		return -1
	}
}

var UInt32Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(uint32) - b.(uint32)
	switch {
	case cmp == 0:
		return 0
	case cmp > 0:
		return 1
	default:
		return -1
	}
}

var UInt64Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(uint64) - b.(uint64)
	switch {
	case cmp == 0:
		return 0
	case cmp > 0:
		return 1
	default:
		return -1
	}
}

var UInt8Comparator Comparator = func(a interface{}, b interface{}) int {
	cmp := a.(uint8) - b.(uint8)
	switch {
	case cmp == 0:
		return 0
	case cmp > 0:
		return 1
	default:
		return -1
	}
}
