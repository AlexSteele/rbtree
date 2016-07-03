package treeset

import (
	"strings"
)

type Comparator func(a interface{}, b interface{}) int

var IntComparator Comparator = func(a interface{}, b interface{}) int {
	return a.(int) - b.(int)
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

var StringComparator Comparator = func(a interface{}, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
}
