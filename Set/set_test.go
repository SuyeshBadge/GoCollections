package set

import (
	"testing"
)

func TestPowerSet(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Get the power set
	powerSet := s.PowerSet()
	powerSet2 := PowerSet(s)

	// Test the length of the power set
	expectedLength := 1 << s.Len() // 2^n
	if len(powerSet) != expectedLength {
		t.Errorf("Expected power set length to be %d, got %d", expectedLength, len(powerSet))
	}

	expectedLength2 := 1 << s.Len() // 2^n
	if len(powerSet2) != expectedLength2 {
		t.Errorf("Expected power set length to be %d, got %d", expectedLength2, len(powerSet2))
	}

	// Test if the power set contains the original set
	originalSetFound := false
	for _, subset := range powerSet {
		if subset.Equal(s) {
			originalSetFound = true
			break
		}
	}
	if !originalSetFound {
		t.Errorf("Expected power set to contain the original set")
	}

	// Test if each subset is a valid subset of the original set
	for _, subset := range powerSet {
		if !subset.IsSubset(s) {
			t.Errorf("Expected subset %v to be a valid subset of the original set", subset)
		}
	}
}

func TestIsProperSuperset(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Create another set
	s2 := NewSet()

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)

	// Test if s is a proper superset of s2
	expectedResult := true
	result := s.IsProperSuperset(s2)
	result2 := IsProperSuperset(s, s2)
	if result != expectedResult {
		t.Errorf("Expected IsProperSuperset to return %v, got %v", expectedResult, result)
	}
	if result2 != expectedResult {
		t.Errorf("Expected IsProperSuperset to return %v, got %v", expectedResult, result2)
	}

	// Create another set
	s3 := NewSet()

	// Add some elements to the third set
	s3.Add(1)
	s3.Add(2)
	s3.Add(3)

	// Test if s3 is a proper superset of s
	expectedResult = false
	result = s3.IsProperSuperset(s)
	result2 = IsProperSuperset(s3, s)
	if result != expectedResult {
		t.Errorf("Expected IsProperSuperset to return %v, got %v", expectedResult, result)
	}
	if result2 != expectedResult {
		t.Errorf("Expected IsProperSuperset to return %v, got %v", expectedResult, result2)
	}
}
func TestIsProperSubset(t *testing.T) {
	// Create two new instances of the Set struct
	s1 := NewSet()
	s2 := NewSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Test if s1 is a proper subset of s2
	expectedResult := true
	result := s1.IsProperSubset(s2)
	result2 := IsProperSubset(s1, s2)
	if result != expectedResult {
		t.Errorf("Expected IsProperSubset to return %v, got %v", expectedResult, result)
	}
	if result2 != expectedResult {
		t.Errorf("Expected IsProperSubset to return %v, got %v", expectedResult, result2)
	}

	// Test if s2 is a proper subset of s1
	expectedResult = false
	result = s2.IsProperSubset(s1)
	result2 = IsProperSubset(s2, s1)
	if result != expectedResult {
		t.Errorf("Expected IsProperSubset to return %v, got %v", expectedResult, result)
	}
	if result2 != expectedResult {
		t.Errorf("Expected IsProperSubset to return %v, got %v", expectedResult, result2)
	}

	// Test if s1 is a proper subset of itself
	expectedResult = false
	result = s1.IsProperSubset(s1)
	if result != expectedResult {
		t.Errorf("Expected IsProperSubset to return %v, got %v", expectedResult, result)
	}
}
func TestClone(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Clone the set
	clone := s.Clone()

	// Test if the clone is equal to the original set
	if !clone.Equal(s) {
		t.Errorf("Expected clone to be equal to the original set")
	}

	// Test if the clone is a separate instance
	if clone == s {
		t.Errorf("Expected clone to be a separate instance from the original set")
	}

	// Test if modifying the clone does not affect the original set
	clone.Add(4)
	if s.Contains(4) {
		t.Errorf("Expected modifying the clone to not affect the original set")
	}
}
func TestIsDisjoint(t *testing.T) {
	// Create two new instances of the Set struct
	s1 := NewSet()
	s2 := NewSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(4)
	s2.Add(5)
	s2.Add(6)

	// Test if s1 and s2 are disjoint
	expectedResult := true
	result := s1.IsDisjoint(s2)
	result2 := IsDisjoint(s1, s2)
	if result != expectedResult {
		t.Errorf("Expected IsDisjoint to return %v, got %v", expectedResult, result)
	}
	if result2 != expectedResult {
		t.Errorf("Expected IsDisjoint to return %v, got %v", expectedResult, result2)
	}

	// Add an element to the second set that is also in the first set
	s2.Add(2)

	// Test if s1 and s2 are disjoint
	expectedResult = false
	result = s1.IsDisjoint(s2)
	result2 = IsDisjoint(s1, s2)
	if result != expectedResult {
		t.Errorf("Expected IsDisjoint to return %v, got %v", expectedResult, result)
	}
	if result2 != expectedResult {
		t.Errorf("Expected IsDisjoint to return %v, got %v", expectedResult, result2)
	}
}
func TestIsSuperset(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Create another set
	s2 := NewSet()

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)

	// Test if s is a superset of s2
	expectedResult := true
	result := s.IsSuperset(s2)
	result2 := IsSuperset(s, s2)
	if result != expectedResult {
		t.Errorf("Expected IsSuperset to return %v, got %v", expectedResult, result)
	}

	if result2 != expectedResult {
		t.Errorf("Expected IsSuperset to return %v, got %v", expectedResult, result2)
	}

	// Create another set
	s3 := NewSet()

	// Add some elements to the third set
	s3.Add(1)
	s3.Add(2)

	// Test if s3 is a superset of s
	expectedResult = false
	result = s3.IsSuperset(s)
	if result != expectedResult {
		t.Errorf("Expected IsSuperset to return %v, got %v", expectedResult, result)
	}
}

func TestIsSubset(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Create another set
	s2 := NewSet()

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	// Test if s2 is a subset of s
	expectedResult := true
	result := s.IsSubset(s2)
	result2 := IsSubset(s, s2)
	if result != expectedResult {
		t.Errorf("Expected IsSubset to return %v, got %v", expectedResult, result)
	}
	if result2 != expectedResult {
		t.Errorf("Expected IsSubset to return %v, got %v", expectedResult, result2)
	}

	// Create another set
	s3 := NewSet()

	// Add some elements to the third set
	s3.Add(1)
	s3.Add(2)

	// Test if s3 is a subset of s
	expectedResult = false
	result = s.IsSubset(s3)
	result2 = IsSubset(s, s3)
	if result != expectedResult {
		t.Errorf("Expected IsSubset to return %v, got %v", expectedResult, result)
	}
	if result2 != expectedResult {
		t.Errorf("Expected IsSubset to return %v, got %v", expectedResult, result2)
	}
}

func TestSymmetricDifference(t *testing.T) {
	// Create two new instances of the Set struct
	s1 := NewSet()
	s2 := NewSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Get the symmetric difference of the two sets
	symDiff := s1.SymmetricDifference(s2)

	// Test if the symmetric difference contains the correct elements
	expectedElements := []interface{}{1, 4}
	for _, element := range expectedElements {
		if !symDiff.Contains(element) {
			t.Errorf("Expected symmetric difference to contain element %v", element)
		}
	}
}

func TestDifference(t *testing.T) {
	// Create two new instances of the Set struct
	s1 := NewSet()
	s2 := NewSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Get the difference of the two sets
	diff := s1.Difference(s2)

	// Test if the difference contains the correct elements
	expectedElements := []interface{}{1}
	for _, element := range expectedElements {
		if !diff.Contains(element) {
			t.Errorf("Expected difference to contain element %v", element)
		}
	}
}

func TestIntersection(t *testing.T) {
	// Create two new instances of the Set struct
	s1 := NewSet()
	s2 := NewSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Get the intersection of the two sets
	intersection := s1.Intersection(s2)

	// Test if the intersection contains the correct elements
	expectedElements := []interface{}{2, 3}
	for _, element := range expectedElements {
		if !intersection.Contains(element) {
			t.Errorf("Expected intersection to contain element %v", element)
		}
	}
}

func TestUnion(t *testing.T) {
	// Create two new instances of the Set struct
	s1 := NewSet()
	s2 := NewSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	// Get the union of the two sets
	union := s1.Union(s2)

	// Test if the union contains the correct elements
	expectedElements := []interface{}{1, 2, 3, 4}
	for _, element := range expectedElements {
		if !union.Contains(element) {
			t.Errorf("Expected union to contain element %v", element)
		}
	}
}

func TestRemove(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Remove an element from the set
	s.Remove(2)

	// Test if the set contains the correct elements
	expectedElements := []interface{}{1, 3}
	for _, element := range expectedElements {
		if !s.Contains(element) {
			t.Errorf("Expected set to contain element %v", element)
		}
	}
}

func TestContains(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Test if the set contains the correct elements
	expectedElements := []interface{}{1, 2, 3}
	for _, element := range expectedElements {
		if !s.Contains(element) {
			t.Errorf("Expected set to contain element %v", element)
		}
	}
}

func TestLen(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Test the length of the set
	expectedLength := 3
	if s.Len() != expectedLength {
		t.Errorf("Expected set length to be %d, got %d", expectedLength, s.Len())
	}
}

func TestClear(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Clear the set
	s.Clear()

	// Test if the set is empty
	if s.Len() != 0 {
		t.Errorf("Expected set to be empty")
	}
}

func TestEqual(t *testing.T) {
	// Create two new instances of the Set struct
	s1 := NewSet()
	s2 := NewSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	// Test if the sets are equal
	if !s1.Equal(s2) {
		t.Errorf("Expected sets to be equal")
	}
	if !Equal(s1, s2) {
		t.Errorf("Expected sets to be equal")
	}

	// Create another set
	s3 := NewSet()

	// Add some elements to the third set
	s3.Add(1)
	s3.Add(2)
	s3.Add(4)
	// Test if the sets are not equal
	if s1.Equal(s3) {
		t.Errorf("Expected sets to not be equal")
	}
	if Equal(s1, s3) {
		t.Errorf("Expected sets to not be equal")
	}

}

func TestIsEmpty(t *testing.T) {

	// Create a new instance of the Set struct
	s := NewSet()

	// Test if the set is empty
	if !s.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Test if the set is empty
	if s.IsEmpty() {
		t.Errorf("Expected set to not be empty")
	}
}

func TestString(t *testing.T) {
	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add("Hello")

	// Test the string representation of the set
	expectedString := "[Hello]"
	if s.String() != expectedString {
		t.Errorf("Expected set string to be %s, got %s", expectedString, s.String())
	}
}

func TestToSlice(t *testing.T) {

	// Create a new instance of the Set struct
	s := NewSet()

	// Add some elements to the set
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Test the slice representation of the set
	expectedSlice := []interface{}{1, 2, 3}
	for _, element := range expectedSlice {
		if !s.Contains(element) {
			t.Errorf("Expected set to contain element %v", element)
		}
	}
}

func TestNewSetFromSlice(t *testing.T) {
	// Create a new set from a slice
	s := NewSetFromSlice([]interface{}{1, 2, 3})

	// Test if the set contains the correct elements
	expectedElements := []interface{}{1, 2, 3}
	for _, element := range expectedElements {
		if !s.Contains(element) {
			t.Errorf("Expected set to contain element %v", element)
		}
	}
}

func TestNewSet(t *testing.T) {
	// Create a new set
	s := NewSet()

	// Test if the set is empty
	if !s.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}
}

func TestCartesianProduct(t *testing.T) {
	// Create two new instances of the Set struct
	s1 := NewSet()
	s2 := NewSet()

	// Add some elements to the first set
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Add some elements to the second set
	s2.Add("a")
	s2.Add("b")

	// Calculate the cartesian product of s1 and s2
	cartesianProduct := s1.CartesianProduct(s2)
	cartesianProduct2 := CartesianProduct(s1, s2)

	// Test the length of the cartesian product
	expectedLength := s1.Len() * s2.Len()
	if len(cartesianProduct) != expectedLength {
		t.Errorf("Expected cartesian product length to be %d, got %d", expectedLength, len(cartesianProduct))
	}

	expectedLength2 := s1.Len() * s2.Len()
	if len(cartesianProduct2) != expectedLength2 {
		t.Errorf("Expected cartesian product length to be %d, got %d", expectedLength2, len(cartesianProduct2))
	}

	// Test if each set in the cartesian product contains the expected elements
}
