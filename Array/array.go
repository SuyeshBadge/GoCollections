package array

import "errors"

// Array represents a collection of values.
type Array struct {
	values   []interface{} // The underlying slice to store the values.
	isStatic bool          // Indicates whether the array is static or dynamic.
}

// NewStaticArray creates a new static array with the specified size.
// The static array is initialized with zero values for each element.
// The size parameter specifies the number of elements in the array.
// Returns a pointer to the newly created static array.
func NewStaticArray(size int) *Array {
	return &Array{
		values:   make([]interface{}, size),
		isStatic: true,
	}
}

// NewDynamicArray creates a new dynamic array.
// It returns a pointer to an Array struct with isStatic set to false.
func NewDynamicArray() *Array {
	return &Array{
		isStatic: false,
	}
}

// Push adds a new element to the end of the array.
// If the array is static and already at its maximum capacity, it returns an error.
func (a *Array) Push(value interface{}) error {
	if a.isStatic && len(a.values) >= cap(a.values) {
		return errors.New("Array capacity is full")
	}
	a.values = append(a.values, value)
	return nil
}

// Pop removes and returns the last element from the array.
// It modifies the underlying array by reducing its length by 1.
// Returns the removed element.
func (a *Array) Pop() interface{} {
	lastElement := a.values[len(a.values)-1]
	a.values = a.values[:len(a.values)-1]
	return lastElement
}

// ToArray returns the underlying array as a slice of interfaces.
func (a *Array) ToArray() []interface{} {
	return a.values
}

// Len returns the length of the array.
func (a *Array) Len() int {
	return len(a.values)
}

// IsEmpty checks if the array is empty.
// It returns true if the array is empty, otherwise false.
func (a *Array) IsEmpty() bool {
	return len(a.values) == 0
}

// Contains checks if the array contains the specified value.
// It returns true if the value is found, otherwise it returns false.
func (a *Array) Contains(value interface{}) bool {
	for _, v := range a.values {
		if v == value {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of the specified value in the array.
// If the value is not found, it returns -1.
func (a *Array) IndexOf(value interface{}) int {
	for i, v := range a.values {
		if v == value {
			return i
		}
	}
	return -1
}

// InsertAt inserts a value at the specified index in the array.
// It returns an error if the index is out of range.
// The index should be a non-negative integer less than or equal to the length of the array.
// The value parameter represents the value to be inserted.
func (a *Array) InsertAt(index int, value interface{}) error {
	if index < 0 || index > len(a.values) {
		return errors.New("index out of range")
	}
	a.values = append(a.values[:index], append([]interface{}{value}, a.values[index:]...)...)
	return nil
}

// RemoveAt removes the element at the specified index from the array.
// It returns an error if the index is out of range.
// The elements after the removed element are shifted to fill the gap.
func (a *Array) RemoveAt(index int) error {
	if index < 0 || index >= len(a.values) {
		return errors.New("index out of range")
	}
	a.values = append(a.values[:index], a.values[index+1:]...)
	return nil
}

// Clear removes all elements from the array.
func (a *Array) Clear() {
	a.values = a.values[:0]
}

// Capacity returns the capacity of the array.
func (a *Array) Capacity() int {
	return cap(a.values)
}

// Resize resizes the array to the specified newSize.
// It creates a new slice with the given newSize and copies the existing values into it.
// If newSize is smaller than the current size, the values beyond newSize will be truncated.
// If newSize is larger than the current size, the additional elements will be initialized with their zero values.
func (a *Array) Resize(newSize int) {
	newSlice := make([]interface{}, newSize)
	copy(newSlice, a.values)
	a.values = newSlice
}

// Get returns the element at the specified index in the array.
// It returns an error if the index is out of range.
func (a *Array) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(a.values) {
		return nil, errors.New("index out of range")
	}
	return a.values[index], nil
}

// Set sets the value at the specified index in the array.
// It returns an error if the index is out of range.
func (a *Array) Set(index int, value interface{}) error {
	if index < 0 || index >= len(a.values) {
		return errors.New("index out of range")
	}
	a.values[index] = value
	return nil
}
