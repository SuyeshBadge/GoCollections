package set

import "fmt"

type Set struct {
	elements map[interface{}]bool
}

// Add adds an item to the set.
// The item parameter is the value to be added to the set.
func (s *Set) Add(item interface{}) {
	s.elements[item] = true
}

// Remove removes the specified item from the set.
func (s *Set) Remove(item interface{}) {
	delete(s.elements, item)
}

// Contains checks if the Set contains the specified item.
// It returns true if the item is found in the Set, otherwise it returns false.
func (s *Set) Contains(item interface{}) bool {
	return s.elements[item]
}

// Len returns the number of elements in the set.
// It returns an integer representing the size of the set.
func (s *Set) Len() int {
	return len(s.elements)
}

// Clear removes all elements from the set.
func (s *Set) Clear() {
	s.elements = make(map[interface{}]bool)
}

// Equal checks if the current set is equal to another set.
// It returns true if the sets have the same length and contain the same elements,
// otherwise it returns false.
func (s *Set) Equal(other *Set) bool {
	if s.Len() != other.Len() {
		return false
	}
	for key := range s.elements {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// ToSlice returns a slice containing all the elements in the set.
// The order of the elements in the slice is not guaranteed.
func (s *Set) ToSlice() []interface{} {
	slice := make([]interface{}, 0, s.Len())
	for key := range s.elements {
		slice = append(slice, key)
	}
	return slice
}

// FromSlice adds all the elements from the given slice to the set.
// It iterates over the slice and calls the Add method to add each element to the set.
func (s *Set) FromSlice(slice []interface{}) {
	for _, item := range slice {
		s.Add(item)
	}
}

// NewSet creates and returns a new Set.
func NewSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

// NewSetFromSlice creates a new Set from a given slice of interface{} values.
// It takes a slice as input and returns a pointer to the newly created Set.
// The elements of the slice are added to the Set using the FromSlice method.
func NewSetFromSlice(slice []interface{}) *Set {
	set := NewSet()
	set.FromSlice(slice)
	return set
}

// Union returns a new Set that contains all the elements from both s1 and s2.
// The original sets s1 and s2 are not modified.
func Union(s1, s2 *Set) *Set {
	s := NewSet()
	for key := range s1.elements {
		s.Add(key)
	}
	for key := range s2.elements {
		s.Add(key)
	}
	return s
}

// Intersection returns a new Set that contains the intersection of two Sets, s1 and s2.
// It iterates over the elements in s1 and checks if each element is present in s2.
// If an element is found in both s1 and s2, it is added to the new Set.
// The resulting Set is then returned.
func Intersection(s1, s2 *Set) *Set {
	s := NewSet()
	for key := range s1.elements {
		if s2.Contains(key) {
			s.Add(key)
		}
	}
	return s
}

// Difference returns a new Set that contains the elements present in s1 but not in s2.
// The function iterates over the elements in s1 and checks if each element is present in s2.
// If an element is not found in s2, it is added to the new Set.
// The returned Set does not modify the original Sets s1 and s2.
func Difference(s1, s2 *Set) *Set {
	s := NewSet()
	for key := range s1.elements {
		if !s2.Contains(key) {
			s.Add(key)
		}
	}
	return s
}

// SymmetricDifference returns a new Set that contains elements that are present in either s1 or s2, but not in both.
// The function takes two pointers to Set objects, s1 and s2, and returns a pointer to a new Set object.
// The returned Set object contains elements that are present in either s1 or s2, but not in both.
// The function iterates over the elements of s1 and checks if each element is present in s2.
// If an element is not present in s2, it is added to the new Set object.
// Then, the function iterates over the elements of s2 and checks if each element is present in s1.
// If an element is not present in s1, it is added to the new Set object.
// Finally, the function returns the new Set object.
func SymmetricDifference(s1, s2 *Set) *Set {
	s := NewSet()
	for key := range s1.elements {
		if !s2.Contains(key) {
			s.Add(key)
		}
	}
	for key := range s2.elements {
		if !s1.Contains(key) {
			s.Add(key)
		}
	}
	return s
}

// IsSubset checks if s1 is a subset of s2.
// It iterates over the elements in s1 and checks if each element is present in s2.
// If any element in s1 is not found in s2, it returns false.
// If all elements in s1 are found in s2, it returns true.
func IsSubset(s1, s2 *Set) bool {
	for key := range s1.elements {
		if !s2.Contains(key) {
			return false
		}
	}
	return true
}

// IsSuperset checks if s1 is a superset of s2.
// It returns true if s1 contains all the elements of s2, false otherwise.
func IsSuperset(s1, s2 *Set) bool {
	return IsSubset(s2, s1)
}

// IsDisjoint checks if two sets are disjoint.
// It returns true if there are no common elements between the two sets, otherwise it returns false.
func IsDisjoint(s1, s2 *Set) bool {
	for key := range s1.elements {
		if s2.Contains(key) {
			return false
		}
	}
	return true
}

// Clone creates a new Set that is a copy of the given Set.
// It returns a pointer to the new Set.
func Clone(s *Set) *Set {
	return NewSetFromSlice(s.ToSlice())
}

// Equal checks if two sets are equal.
// It returns true if the sets are equal, and false otherwise.
func Equal(s1, s2 *Set) bool {
	return s1.Equal(s2)
}

// String returns a string representation of the Set.
// It converts the Set into a slice and formats it using the fmt.Sprintf function.
func (s *Set) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}

// IsEmpty checks if the Set is empty.
// It returns true if the Set is empty, otherwise it returns false.
func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

// IsProperSubset checks if s1 is a proper subset of s2.
// It returns true if s1 is a proper subset of s2, otherwise it returns false.
func IsProperSubset(s1, s2 *Set) bool {
	return s1.Len() < s2.Len() && IsSubset(s1, s2)
}

// IsProperSuperset checks if s1 is a proper superset of s2.
// It returns true if s1 is a proper superset of s2, otherwise it returns false.
func IsProperSuperset(s1, s2 *Set) bool {
	return s1.Len() > s2.Len() && IsSuperset(s1, s2)
}

// PowerSet returns the power set of the given Set.
// It returns a slice of Sets, where each Set is a subset of the original Set.
func PowerSet(s *Set) []*Set {
	powerSet := make([]*Set, 0)
	powerSet = append(powerSet, NewSet())
	for _, item := range s.ToSlice() {
		for _, subset := range powerSet {
			newSubset := Clone(subset)
			newSubset.Add(item)
			powerSet = append(powerSet, newSubset)
		}
	}
	return powerSet
}

// CartesianProduct returns the Cartesian product of two Sets.
// It returns a slice of Sets, where each Set is a pair of elements from the two Sets.
func CartesianProduct(s1, s2 *Set) []*Set {
	cartesianProduct := make([]*Set, 0)
	for _, item1 := range s1.ToSlice() {
		for _, item2 := range s2.ToSlice() {
			cartesianProduct = append(cartesianProduct, NewSetFromSlice([]interface{}{item1, item2}))
		}
	}
	return cartesianProduct
}

// Difference returns a new Set that contains the elements that are present in the receiver Set but not in the given Set s2.
// The receiver Set remains unchanged.
func (s *Set) Difference(s2 *Set) *Set {
	return Difference(s, s2)
}

// SymmetricDifference returns a new Set that contains the elements that are present in either the current Set or the given Set, but not in both.
// The original Sets are not modified.
func (s *Set) SymmetricDifference(s2 *Set) *Set {
	return SymmetricDifference(s, s2)
}

// Union returns a new Set that contains all the elements from both the current Set and the given Set.
// The original Sets are not modified.
func (s *Set) Union(s2 *Set) *Set {
	return Union(s, s2)
}

// IsSubset checks if the current Set is a subset of the given Set s2.
func (s *Set) IsSubset(s2 *Set) bool {
	for element := range s.elements {
		if !s2.Contains(element) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the current Set is a superset of the given Set s2.
func (s *Set) IsSuperset(s2 *Set) bool {
	return s2.IsSubset(s)
}

// IsDisjoint checks if the current Set and the given Set s2 are disjoint.
func (s *Set) IsDisjoint(s2 *Set) bool {
	for element := range s.elements {
		if s2.Contains(element) {
			return false
		}
	}
	return true
}

// Clone creates a new Set that is a copy of the current Set.
func (s *Set) Clone() *Set {
	clone := NewSet()
	for element := range s.elements {
		clone.Add(element)
	}
	return clone
}

// IsProperSubset checks if the current Set is a proper subset of the given Set s2.
func (s *Set) IsProperSubset(s2 *Set) bool {
	return s.IsSubset(s2) && !s.Equal(s2)
}

// IsProperSuperset checks if the current Set is a proper superset of the given Set s2.
func (s *Set) IsProperSuperset(s2 *Set) bool {
	return s.IsSuperset(s2) && !s.Equal(s2)
}

// PowerSet returns the power set of the current Set.
func (s *Set) PowerSet() []*Set {
	powerSet := make([]*Set, 0)
	powerSet = append(powerSet, NewSet())
	for _, item := range s.ToSlice() {
		for _, subset := range powerSet {
			newSubset := subset.Clone()
			newSubset.Add(item)
			powerSet = append(powerSet, newSubset)
		}
	}
	return powerSet
}

// CartesianProduct returns the Cartesian product of the current Set and the given Set s2.
func (s *Set) CartesianProduct(s2 *Set) []*Set {
	cartesianProduct := make([]*Set, 0)
	for _, item1 := range s.ToSlice() {
		for _, item2 := range s2.ToSlice() {
			cartesianProduct = append(cartesianProduct, NewSetFromSlice([]interface{}{item1, item2}))
		}
	}
	return cartesianProduct
}

// Intersection returns a new Set that contains the intersection of the current Set and the given Set s2.
func (s *Set) Intersection(s2 *Set) *Set {
	return Intersection(s, s2)
}
