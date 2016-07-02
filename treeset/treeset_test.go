package treeset

import (
	"testing"
)

func TestAdd_SortedOrder(t *testing.T) {
	s := New(IntComparator)
	elems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range elems {
		res := s.Add(v)
		if res != nil {
			t.Fatalf("Unexpected return: %v", res)
		}
	}
	for _, v := range elems {
		if !s.Contains(v) {
			t.Fatalf("Set did not contain %v", v)
		}
	}
}

func TestAdd_UnsortedOrder(t *testing.T) {
	
	// Admittedly this test could be more precise.
	s := New(IntComparator)
	elems := []int{80, 15, 30, 10, 1, 2, 90, 7, 23, 26, 83}
	for _, v := range elems {
		res := s.Add(v)
		if res != nil {
			t.Fatalf("Unexpected return: %v", res)
		}
	}
	for _, v := range elems {
		if !s.Contains(v) {
			t.Fatalf("Set did not contain %v", v)
		}
	}
}

func TestAdd_Alternating(t *testing.T) {
	s := New(IntComparator)
	elems := []int{100, 50, 150, 75, 125, 25, 175, 40, 160, 10, 190}
	for _, v := range elems {
		res := s.Add(v)
		if res != nil {
			t.Fatalf("Unexpected return: %v", res)
		}
	}
	for _, v := range elems {
		if !s.Contains(v) {
			t.Fatalf("Set did not contain %v", v)
		}
	}
}

func TestAdd_LeftTree(t *testing.T) {
	s := New(IntComparator)
	elems := []int{100, 50, 75, 25, 20, 40, 15, 30, 10, 5}
	for _, v := range elems {
		res := s.Add(v)
		if res != nil {
			t.Fatalf("Unexpected return: %v", res)
		}
	}
	for _, v := range elems {
		if !s.Contains(v) {
			t.Fatalf("Set did not contain %v", v)
		}
	}
}

func TestAdd_RightTree(t *testing.T) {
	s := New(IntComparator)
	elems := []int{100, 150, 125, 175, 160, 180, 140, 155, 190}
	for _, v := range elems {
		res := s.Add(v)
		if res != nil {
			t.Fatalf("Unexpected return: %v", res)
		}
	}
	for _, v := range elems {
		if !s.Contains(v) {
			t.Fatalf("Set did not contain %v", v)
		}
	}
}

func TestAdd_DuplicateElement(t *testing.T) {
	s := New(IntComparator)
	s.Add(1)
	old := s.Add(1)
	if old != 1 {
		t.Fatal("Add duplicate did not return old element.")
	}
	if s.Length() != 1 {
		t.Fatal("Add duplicate changed length.")
	}
}

func TestRemove_Root(t *testing.T) {
	s := New(StringComparator)
	s.Add("abc")
	if !s.Remove("abc") {
		t.Fatal("Failed to remove root.")
	}
	if s.Length() != 0 {
		t.Fatal("Removal of root didn't set length.")
	}
	if s.Contains("abc") {
		t.Fatal("Set still contains removed element.")
	}
}

func TestRemove_NonRootLeaf(t *testing.T) {
	s := New(StringComparator)
	s.Add("a")
	s.Add("c")
	s.Add("d")
	if !s.Remove("d") {
		t.Fatal("Failed to remove non-root leaf.")
	}
	if s.Length() != 2 {
		t.Fatal("Removal of element didn't set length.")
	}
	if s.Contains("d") {
		t.Fatal("Set still contains removed element.")
	}

	s.Add("b")
	if !s.Remove("b") {
		t.Fatal("Failed to remove non-root leaf.")
	}
	if s.Length() != 2 {
		t.Fatal("Removal of element didn't set length.")
	}
	if s.Contains("b") {
		t.Fatal("Set still contains removed element.")
	}
}

func TestRemove_Bulk(t *testing.T) {
	s := New(IntComparator)
	elems := []int{9, 7, 17, 2, 20, 6, 10, 3, 11}
	for _, v := range elems {
		s.Add(v)
	}
	for _, v := range elems {
		if ! s.Remove(v) {
			t.Fatalf("Failed to remove %v", v)
		}
	}
	
}

func TestContains(t *testing.T) {
	
}

func TestClear(t *testing.T) {
	s := New(IntComparator)
	elems := []int{5, 4, 3, 2, 1} 
	for _, v := range elems {
		s.Add(v)
	}

	s.Clear()

	if s.Length() != 0 {
		t.Fatal("Set had nonzero length.")
	}

	for _, v := range elems {
		if s.Contains(v) {
			t.Fatalf("Set contained %v after clear.", v)
		}
	}
}

func TestLength(t *testing.T) {
	t.Skip() // TODO: REMOVE
	s := New(IntComparator)
	elems := []int{19, 73, 930, 1694, 3910, 82, 17, 16, 15, 14, 91}
	if s.Length() != 0 {
		t.Fatal("Nonzero length.")
	}
	for k, v := range elems {
		s.Add(v)
		if l := s.Length(); l != len(elems[:k+1]) {
			t.Fatalf("Incorrect length. Expected %v. Got %v", len(elems[:k+1]), l)
		}
	}
	for k, v := range elems {
		s.Remove(v)
		if l := s.Length(); l != len(elems) - len(elems[:k+1]) {
			t.Fatalf("Incorrect length. Expected %v. Got %v", len(elems) - len(elems[:k+1]), l)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	
}

// Benchmarks.

func BenchmarkAdd(t *testing.B) {
	
}

func BenchmarkRemove(t *testing.B) {
	
}
