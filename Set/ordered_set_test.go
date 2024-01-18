package set

import (
	"testing"
)

func TestOrderedSetCartesianProduct(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(4)
	s2.Add(5)

	// Calculate the cartesian product of the two sets
	cartesianProduct := s1.CartesianProduct(s2)

	// Test the length of the cartesian product
	expectedLength := s1.Len() * s2.Len()
	if len(cartesianProduct) != expectedLength {
		t.Errorf("Expected cartesian product length to be %d, got %d", expectedLength, len(cartesianProduct))
	}

	// Test if each element in the cartesian product is a valid combination
	for _, product := range cartesianProduct {
		if product.Len() != 2 {
			t.Errorf("Expected cartesian product element length to be 2, got %d", product.Len())
		}
		if !s1.Contains(product.Get(0)) {
			t.Errorf("Expected cartesian product element to contain element from first set")
		}
		if !s2.Contains(product.Get(1)) {
			t.Errorf("Expected cartesian product element to contain element from second set")
		}
	}
}

func TestOrderedSetDifference(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Calculate the difference of the two sets
	difference := s1.Difference(s2)

	// Test the length of the difference
	expectedLength := 1
	if difference.Len() != expectedLength {
		t.Errorf("Expected difference length to be %d, got %d", expectedLength, difference.Len())
	}

	// Test if the difference contains the correct elements
	if !difference.Contains(1) {
		t.Errorf("Expected difference to contain element 1")
	}
}

func TestOrderedSetIntersection(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Calculate the intersection of the two sets
	intersection := s1.Intersection(s2)

	// Test the length of the intersection
	expectedLength := 2
	if intersection.Len() != expectedLength {
		t.Errorf("Expected intersection length to be %d, got %d", expectedLength, intersection.Len())
	}

	// Test if the intersection contains the correct elements
	if !intersection.Contains(2) {
		t.Errorf("Expected intersection to contain element 2")
	}
	if !intersection.Contains(3) {
		t.Errorf("Expected intersection to contain element 3")
	}
}

func TestOrderedSetIsSuperset(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s1.Add(4)
	s1.Add(5)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	// Test if the first set is a superset of the second set
	if !s1.IsSuperset(s2) {
		t.Errorf("Expected first set to be a superset of the second set")
	}
}

func TestOrderedSetSymmetricDifference(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Calculate the symmetric difference of the two sets
	symDiff := s1.SymmetricDifference(s2)

	// Test the length of the symmetric difference
	expectedLength := 2
	if symDiff.Len() != expectedLength {
		t.Errorf("Expected symmetric difference length to be %d, got %d", expectedLength, symDiff.Len())
	}

	// Test if the symmetric difference contains the correct elements
	if !symDiff.Contains(1) {
		t.Errorf("Expected symmetric difference to contain element 1")
	}
	if !symDiff.Contains(4) {
		t.Errorf("Expected symmetric difference to contain element 4")
	}
}

func TestOrderedSetUnion(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Calculate the union of the two sets
	union := s1.Union(s2)

	// Test the length of the union
	expectedLength := 4
	if union.Len() != expectedLength {
		t.Errorf("Expected union length to be %d, got %d", expectedLength, union.Len())
	}

	// Test if the union contains the correct elements
	if !union.Contains(1) {
		t.Errorf("Expected union to contain element 1")
	}
	if !union.Contains(2) {
		t.Errorf("Expected union to contain element 2")
	}
	if !union.Contains(3) {
		t.Errorf("Expected union to contain element 3")
	}
	if !union.Contains(4) {
		t.Errorf("Expected union to contain element 4")
	}
}

func TestOrderedSetToSlice(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Convert the set to a slice
	slice := s.ToSlice()

	// Test the length of the slice
	expectedLength := 3
	if len(slice) != expectedLength {
		t.Errorf("Expected slice length to be %d, got %d", expectedLength, len(slice))
	}

	// Test if the slice contains the correct elements
	if slice[0] != 1 {
		t.Errorf("Expected slice to contain element 1")
	}
	if slice[1] != 2 {
		t.Errorf("Expected slice to contain element 2")
	}
	if slice[2] != 3 {
		t.Errorf("Expected slice to contain element 3")
	}
}

func TestOrderedSetString(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Convert the set to a string
	str := s.String()

	// Test the string representation of the set
	expectedString := "[1 2 3]"
	if str != expectedString {
		t.Errorf("Expected string representation of set to be %s, got %s", expectedString, str)
	}
}

func TestOrderedSetClear(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Clear the set
	s.Clear()

	// Test the length of the set
	expectedLength := 0
	if s.Len() != expectedLength {
		t.Errorf("Expected set length to be %d, got %d", expectedLength, s.Len())
	}
}

func TestOrderedSetContains(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Test if the set contains the correct elements
	if !s.Contains(1) {
		t.Errorf("Expected set to contain element 1")
	}
	if !s.Contains(2) {
		t.Errorf("Expected set to contain element 2")
	}
	if !s.Contains(3) {
		t.Errorf("Expected set to contain element 3")
	}
}

func TestOrderedSetCopy(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Create a copy of the set
	copy := s.Copy()

	// Test if the copy contains the correct elements
	if !copy.Contains(1) {
		t.Errorf("Expected copy to contain element 1")
	}
	if !copy.Contains(2) {
		t.Errorf("Expected copy to contain element 2")
	}
	if !copy.Contains(3) {
		t.Errorf("Expected copy to contain element 3")
	}
}

func TestOrderedSetPowerSet(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Calculate the power set of the set
	powerSet := s.PowerSet()

	// Test the length of the power set
	expectedLength := 8
	if len(powerSet) != expectedLength {
		t.Errorf("Expected power set length to be %d, got %d", expectedLength, len(powerSet))
	}

	// Test if each element in the power set is a valid subset
	for _, subset := range powerSet {
		if !subset.IsSubset(s) {
			t.Errorf("Expected power set element to be a subset of the original set")
		}
	}
}

func TestOrderedSetRemove(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add some elements to the set
	s.Add("a")
	s.Add("b")
	s.Add("c")

	// Remove an element from the set
	s.Remove("b")

	// Test the length of the set
	expectedLength := 2
	if s.Len() != expectedLength {
		t.Errorf("Expected set length to be %d, got %d", expectedLength, s.Len())
	}

	// Test if the set contains the correct elements
	if !s.Contains("a") {
		t.Errorf("Expected set to contain element a")
	}
	if s.Contains("b") {
		t.Errorf("Expected set to not contain element b")
	}
	if !s.Contains("c") {
		t.Errorf("Expected set to contain element c")
	}
}

func TestOrderedSetLen(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Test the length of the set
	expectedLength := 0
	if s.Len() != expectedLength {
		t.Errorf("Expected set length to be %d, got %d", expectedLength, s.Len())
	}
}

func TestOrderedSetAdd(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add an element to the set
	s.Add(1)

	// Test the length of the set
	expectedLength := 1
	if s.Len() != expectedLength {
		t.Errorf("Expected set length to be %d, got %d", expectedLength, s.Len())
	}

	// Test if the set contains the correct element
	if !s.Contains(1) {
		t.Errorf("Expected set to contain element 1")
	}
}

func TestOrderedSetGet(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Test getting elements at different indices
	tests := []struct {
		index          int
		expectedResult interface{}
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, nil},  // Index out of range
		{-1, nil}, // Index out of range
	}

	for _, test := range tests {
		result := s.Get(test.index)
		if result != test.expectedResult {
			t.Errorf("Expected element at index %d to be %v, got %v", test.index, test.expectedResult, result)
		}
	}
}
func TestOrderedSetEqual(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	// Test if the two sets are equal
	if !s1.Equal(s2) {
		t.Errorf("Expected sets to be equal")
	}

	// Add an extra element to the second set
	s2.Add(4)

	// Test if the two sets are still equal
	if s1.Equal(s2) {
		t.Errorf("Expected sets to be unequal")
	}
}

func TestOrderedSetIsSubset(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)

	// Test if the first set is a subset of the second set
	if !s1.IsSubset(s2) {
		t.Errorf("Expected first set to be a subset of the second set")
	}

	// Create two new instances of the OrderedSet struct
	s3 := NewOrderedSet()
	s4 := NewOrderedSet()

	// Add some elements to the first set
	s3.Add(1)
	s3.Add(2)
	s3.Add(3)

	// Add some elements to the second set
	s4.Add(4)
	s4.Add(5)

	// Test if the first set is a subset of the second set
	if s3.IsSubset(s4) {
		t.Errorf("Expected first set to not be a subset of the second set")
	}
}

func TestOrderedSetIsDisjoint(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(4)
	s2.Add(5)

	// Test if the two sets are disjoint
	if !s1.IsDisjoint(s2) {
		t.Errorf("Expected sets to be disjoint")
	}

	// Add some elements to make the sets not disjoint
	s2.Add(2)

	// Test if the two sets are disjoint after adding elements
	if s1.IsDisjoint(s2) {
		t.Errorf("Expected sets to not be disjoint")
	}
}

func TestOrderedSetIsProperSuperset(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s1.Add(4)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	// Test if the first set is a proper superset of the second set
	if !s1.IsProperSuperset(s2) {
		t.Errorf("Expected first set to be a proper superset of the second set")
	}

	// Test if the second set is a proper superset of the first set
	if s2.IsProperSuperset(s1) {
		t.Errorf("Expected second set not to be a proper superset of the first set")
	}
}

func TestOrderedSetIsProperSubset(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)

	// Test if the first set is a proper subset of the second set
	if !s1.IsProperSubset(s2) {
		t.Errorf("Expected first set to be a proper subset of the second set")
	}

	// Add an element to the first set
	s1.Add(7)

	// Test if the first set is still a proper subset of the second set
	if s1.IsProperSubset(s2) {
		t.Errorf("Expected first set to not be a proper subset of the second set after adding an element")
	}
}

func TestOrderedSetDisjointUnion(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(4)
	s2.Add(5)

	// Calculate the disjoint union of the two sets
	disjointUnion := s1.DisjointUnion(s2)

	// Test the length of the disjoint union
	expectedLength := 2
	if len(disjointUnion) != expectedLength {
		t.Errorf("Expected disjoint union length to be %d, got %d", expectedLength, len(disjointUnion))
	}

	// Test if the disjoint union contains the correct sets
	if disjointUnion[0] != s1 {
		t.Errorf("Expected disjoint union to contain first set")
	}
	if disjointUnion[1] != s2 {
		t.Errorf("Expected disjoint union to contain second set")
	}
}

func TestOrderedSetIsEmpty(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Test if the set is empty
	if !s.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}

	// Add an element to the set
	s.Add(1)

	// Test if the set is not empty
	if s.IsEmpty() {
		t.Errorf("Expected set to not be empty")
	}
}
func TestOrderedSetPop(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Test the Pop method
	s.Pop()

	// Test if the length of the set has decreased by 1
	expectedLength := 2
	if s.Len() != expectedLength {
		t.Errorf("Expected set length to be %d, got %d", expectedLength, s.Len())
	}
}

func TestOrderedSetPopEmptySet(t *testing.T) {
	// Create a new instance of the OrderedSet struct
	s := NewOrderedSet()

	// Test the Pop method on an empty set
	poppedItem := s.Pop()

	// Test if the popped item is nil
	if poppedItem != nil {
		t.Errorf("Expected popped item to be nil")
	}

	// Test if the length of the set is still 0
	expectedLength := 0
	if s.Len() != expectedLength {
		t.Errorf("Expected set length to be %d, got %d", expectedLength, s.Len())
	}
}
func TestOrderedSetIsEqual(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	// Test if the two sets are equal
	if !s1.IsEqual(s2) {
		t.Errorf("Expected the two sets to be equal")
	}
}

func TestOrderedSetIsEqualDifferentOrder(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set in a different order
	s2.Add(3)
	s2.Add(2)
	s2.Add(1)

	// Test if the two sets are equal
	if !s1.IsEqual(s2) {
		t.Errorf("Expected the two sets to be equal")
	}
}

func TestOrderedSetIsEqualDifferentElements(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some different elements to the second set
	s2.Add(4)
	s2.Add(5)
	s2.Add(6)

	// Test if the two sets are equal
	if s1.IsEqual(s2) {
		t.Errorf("Expected the two sets to be different")
	}
}
func TestOrderedSetIsStrictSubset(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Test if the first set is a strict subset of the second set
	if !s1.IsStrictSubset(s2) {
		t.Errorf("Expected first set to be a strict subset of the second set")
	}
}

func TestOrderedSetIsStrictSubset_NotStrictSubset(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	// Test if the first set is a strict subset of the second set
	if s1.IsStrictSubset(s2) {
		t.Errorf("Expected first set to not be a strict subset of the second set")
	}
}
func TestOrderedSetIsStrictSuperset(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s1.Add(4)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)

	// Test if the first set is a strict superset of the second set
	if !s1.IsStrictSuperset(s2) {
		t.Errorf("Expected first set to be a strict superset of the second set")
	}
}

func TestOrderedSetIsStrictSuperset_EmptySet(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2.Add(1)
	s2.Add(2)

	// Test if the first set is a strict superset of the second set
	if !s1.IsStrictSuperset(s2) {
		t.Errorf("Expected first set to be a strict superset of the second set")
	}
}

func TestOrderedSetIsStrictSuperset_NotSuperset(t *testing.T) {
	// Create two new instances of the OrderedSet struct
	s1 := NewOrderedSet()
	s2 := NewOrderedSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Test if the first set is a strict superset of the second set
	if s1.IsStrictSuperset(s2) {
		t.Errorf("Expected first set to not be a strict superset of the second set")
	}
}
