package treeset

import (
	"strings"

	"github.com/alexsteele/go-sets/set"
)

var IntComparator set.Comparator = func(a interface{}, b interface{}) int {
	return a.(int) - b.(int)
}

var Float64Comparator set.Comparator = func(a interface{}, b interface{}) int {
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

var StringComparator set.Comparator = func(a interface{}, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
}

