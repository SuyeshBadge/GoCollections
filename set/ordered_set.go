package set

import "fmt"

type OrderedSet struct {
	elements []interface{}
}

// NewOrderedSet creates a new OrderedSet.
func NewOrderedSet() *OrderedSet {
	return &OrderedSet{make([]interface{}, 0)}
}

// Add adds an item to the set.
// The item parameter is the value to be added to the set.
func (s *OrderedSet) Add(item interface{}) {
	if !s.Contains(item) {
		s.elements = append(s.elements, item)
	}
}

// Remove removes the specified item from the set.
func (s *OrderedSet) Remove(item interface{}) {
	for i, element := range s.elements {
		if element == item {
			s.elements = append(s.elements[:i], s.elements[i+1:]...)
			break
		}
	}
}

// Get returns the item at the specified index.
// It returns the item at the specified index if the index is valid, otherwise it returns nil.
func (s *OrderedSet) Get(index int) interface{} {
	if index < 0 || index >= s.Len() {
		return nil
	}
	return s.elements[index]
}

// Contains checks if the Set contains the specified item.
// It returns true if the item is found in the Set, otherwise it returns false.
func (s *OrderedSet) Contains(item interface{}) bool {
	for _, element := range s.elements {
		if element == item {
			return true
		}
	}
	return false
}

// Len returns the number of elements in the set.
// It returns an integer representing the size of the set.
func (s *OrderedSet) Len() int {
	return len(s.elements)
}

// Clear removes all elements from the set.
func (s *OrderedSet) Clear() {
	s.elements = make([]interface{}, 0)
}

// Equal checks if the current set is equal to another set.
// It returns true if the sets have the same length and contain the same elements,
// otherwise it returns false.
func (s *OrderedSet) Equal(other *OrderedSet) bool {
	if s.Len() != other.Len() {
		return false
	}
	for _, element := range s.elements {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// ToSlice returns a slice containing all the elements in the set.
// The order of the elements in the slice is guaranteed.
func (s *OrderedSet) ToSlice() []interface{} {
	return s.elements
}

// String returns a string representation of the set.
func (s *OrderedSet) String() string {
	return fmt.Sprintf("%v", s.elements)
}

// Union returns a new set containing all the elements in the current set and the other set.
func (s *OrderedSet) Union(other *OrderedSet) *OrderedSet {
	union := NewOrderedSet()
	for _, element := range s.elements {
		union.Add(element)
	}
	for _, element := range other.elements {
		union.Add(element)
	}
	return union
}

// Intersection returns a new set containing the elements that are in both the current set and the other set.
func (s *OrderedSet) Intersection(other *OrderedSet) *OrderedSet {
	intersection := NewOrderedSet()
	for _, element := range s.elements {
		if other.Contains(element) {
			intersection.Add(element)
		}
	}
	return intersection
}

// Difference returns a new set containing the elements that are in the current set but not in the other set.
func (s *OrderedSet) Difference(other *OrderedSet) *OrderedSet {
	difference := NewOrderedSet()
	for _, element := range s.elements {
		if !other.Contains(element) {
			difference.Add(element)
		}
	}
	return difference
}

// SymmetricDifference returns a new set containing the elements that are in the current set or in the other set but not in both.
func (s *OrderedSet) SymmetricDifference(other *OrderedSet) *OrderedSet {
	symDiff := NewOrderedSet()
	for _, element := range s.elements {
		if !other.Contains(element) {
			symDiff.Add(element)
		}
	}
	for _, element := range other.elements {
		if !s.Contains(element) {
			symDiff.Add(element)
		}
	}
	return symDiff
}

// IsSubset checks if the current set is a subset of the other set.
func (s *OrderedSet) IsSubset(other *OrderedSet) bool {
	for _, element := range s.elements {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the current set is a superset of the other set.
func (s *OrderedSet) IsSuperset(other *OrderedSet) bool {
	return other.IsSubset(s)
}

// IsDisjoint checks if the current set and the other set are disjoint.
func (s *OrderedSet) IsDisjoint(other *OrderedSet) bool {
	for _, element := range s.elements {
		if other.Contains(element) {
			return false
		}
	}
	return true
}

// Clone creates a new set that is a copy of the current set.
func (s *OrderedSet) Clone() *OrderedSet {
	clone := NewOrderedSet()
	for _, element := range s.elements {
		clone.Add(element)
	}
	return clone
}

// IsProperSubset checks if the current set is a proper subset of the other set.
func (s *OrderedSet) IsProperSubset(other *OrderedSet) bool {
	return s.IsSubset(other) && !s.Equal(other)
}

// IsProperSuperset checks if the current set is a proper superset of the other set.
func (s *OrderedSet) IsProperSuperset(other *OrderedSet) bool {
	return s.IsSuperset(other) && !s.Equal(other)
}

// PowerSet returns the power set of the current set.
func (s *OrderedSet) PowerSet() []*OrderedSet {
	powerSet := make([]*OrderedSet, 0)
	powerSet = append(powerSet, NewOrderedSet())
	for _, item := range s.ToSlice() {
		for _, subset := range powerSet {
			newSubset := subset.Clone()
			newSubset.Add(item)
			powerSet = append(powerSet, newSubset)
		}
	}
	return powerSet
}

// CartesianProduct returns the Cartesian product of the current set and the other set.
func (s *OrderedSet) CartesianProduct(other *OrderedSet) []*OrderedSet {
	cartesianProduct := make([]*OrderedSet, 0)
	for _, element := range s.elements {
		for _, element2 := range other.elements {
			product := NewOrderedSet()
			product.Add(element)
			product.Add(element2)
			cartesianProduct = append(cartesianProduct, product)
		}
	}
	return cartesianProduct
}

// DisjointUnion returns the disjoint union of the current set and the other set.
func (s *OrderedSet) DisjointUnion(other *OrderedSet) []*OrderedSet {
	disjointUnion := make([]*OrderedSet, 0)
	disjointUnion = append(disjointUnion, s)
	disjointUnion = append(disjointUnion, other)
	return disjointUnion
}

// IsEmpty checks if the current set is empty.
func (s *OrderedSet) IsEmpty() bool {
	return s.Len() == 0
}

// Pop removes and returns an arbitrary item from the set.
func (s *OrderedSet) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	item := s.elements[0]
	s.Remove(item)
	return item
}

// IsEqual checks if the current set is equal to the other set.
func (s *OrderedSet) IsEqual(other *OrderedSet) bool {
	return s.Equal(other)
}

// IsStrictSubset checks if the current set is a strict subset of the other set.
func (s *OrderedSet) IsStrictSubset(other *OrderedSet) bool {
	return s.IsProperSubset(other)
}

// IsStrictSuperset checks if the current set is a strict superset of the other set.
func (s *OrderedSet) IsStrictSuperset(other *OrderedSet) bool {
	return s.IsProperSuperset(other)
}

// Copy returns a new OrderedSet that is a copy of the original OrderedSet.
// The elements in the new OrderedSet are the same as the elements in the original OrderedSet.
// Changes made to the new OrderedSet will not affect the original OrderedSet.
func (s *OrderedSet) Copy() *OrderedSet {
	return s.Clone()
}
