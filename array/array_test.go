package array

import (
	"testing"
)

func TestInsertAt(t *testing.T) {
	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Test inserting at index 0
	err := arr.InsertAt(0, 10)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val, _ := arr.Get(0); val != 10 {
		t.Errorf("Expected value at index 0 to be 10, got %v", val)
	}

	// Test inserting at index 1
	err = arr.InsertAt(1, 20)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val, _ := arr.Get(1); val != 20 {
		t.Errorf("Expected value at index 1 to be 20, got %v", val)
	}

	// Test inserting at index 2
	err = arr.InsertAt(2, 30)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val, _ := arr.Get(2); val != 30 {
		t.Errorf("Expected value at index 2 to be 30, got %v", val)
	}

	// Test inserting at index out of range
	err = arr.InsertAt(4, 40)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if err.Error() != "index out of range" {
		t.Errorf("Expected error message 'index out of range', got %v", err.Error())
	}
}

func TestRemoveAt(t *testing.T) {
	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Insert some values
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(2, 30)

	// Test removing at index 0
	err := arr.RemoveAt(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val, _ := arr.Get(0); val != 20 {
		t.Errorf("Expected value at index 0 to be 20, got %v", val)
	}

	// Test removing at index 0 again
	err = arr.RemoveAt(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val, _ := arr.Get(0); val != 30 {
		t.Errorf("Expected value at index 0 to be 30, got %v", val)
	}

	// Test removing at index 0 once more
	err = arr.RemoveAt(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val, _ := arr.Get(0); val != nil {
		t.Errorf("Expected value at index 0 to be nil, got %v", val)
	}

	// Test removing at index out of range
	err = arr.RemoveAt(0)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if err.Error() != "index out of range" {
		t.Errorf("Expected error message 'index out of range', got %v", err.Error())
	}
}

func TestGet(t *testing.T) {
	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Test getting value at index 0
	arr.InsertAt(0, 10)
	if val, _ := arr.Get(0); val != 10 {
		t.Errorf("Expected value at index 0 to be 10, got %v", val)
	}

	// Test getting value at index 1
	arr.InsertAt(1, 20)
	if val, _ := arr.Get(1); val != 20 {
		t.Errorf("Expected value at index 1 to be 20, got %v", val)
	}

	// Test getting value at index 2
	arr.InsertAt(2, 30)
	if val, _ := arr.Get(2); val != 30 {
		t.Errorf("Expected value at index 2 to be 30, got %v", val)
	}

	// Test getting value at index out of range
	var err error
	if _, err = arr.Get(3); err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if err.Error() != "index out of range" {
		t.Errorf("Expected error message 'index out of range', got %v", err.Error())
	}
}

func TestSet(t *testing.T) {
	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Test setting value at index 0
	arr.InsertAt(0, 10)
	arr.Set(0, 20)
	if val, _ := arr.Get(0); val != 20 {
		t.Errorf("Expected value at index 0 to be 20, got %v", val)
	}

	// Test setting value at index 1
	arr.InsertAt(1, 30)
	arr.Set(1, 40)
	if val, _ := arr.Get(1); val != 40 {
		t.Errorf("Expected value at index 1 to be 40, got %v", val)
	}

	// Test setting value at index 2
	arr.InsertAt(2, 50)
	arr.Set(2, 60)
	if val, _ := arr.Get(2); val != 60 {
		t.Errorf("Expected value at index 2 to be 60, got %v", val)
	}

	// Test setting value at index out of range
	err := arr.Set(3, 70)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if err.Error() != "index out of range" {
		t.Errorf("Expected error message 'index out of range', got %v", err.Error())
	}
}

func TestClear(t *testing.T) {

	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Insert some values
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(2, 30)

	// Test clearing the array
	arr.Clear()
	if len(arr.ToArray()) != 0 {
		t.Errorf("Expected array length to be 0, got %v", len(arr.ToArray()))
	}
}

func TestResize(t *testing.T) {

	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Insert some values
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(2, 30)

	// Test resizing to a smaller size
	arr.Resize(2)
	if len(arr.ToArray()) != 2 {
		t.Errorf("Expected array length to be 2, got %v", len(arr.ToArray()))
	}

	// Test resizing to a larger size
	arr.Resize(4)
	if len(arr.ToArray()) != 4 {
		t.Errorf("Expected array length to be 4, got %v", len(arr.ToArray()))
	}
	if arr.ToArray()[2] != nil {
		t.Errorf("Expected value at index 2 to be nil, got %v", arr.ToArray()[2])
	}
	if arr.ToArray()[3] != nil {
		t.Errorf("Expected value at index 3 to be nil, got %v", arr.ToArray()[3])
	}
}

func TestIndexOf(t *testing.T) {

	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Insert some values
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(2, 30)

	// Test finding the index of a value
	if index := arr.IndexOf(20); index != 1 {
		t.Errorf("Expected index of 1, got %v", index)
	}

	// Test finding the index of a value that doesn't exist
	if index := arr.IndexOf(40); index != -1 {
		t.Errorf("Expected index of -1, got %v", index)
	}
}

func TestContains(t *testing.T) {

	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Insert some values
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(2, 30)

	// Test finding a value that exists
	if !arr.Contains(20) {
		t.Errorf("Expected Contains to return true, got false")
	}

	// Test finding a value that doesn't exist
	if arr.Contains(40) {
		t.Errorf("Expected Contains to return false, got true")
	}
}

func TestIsEmpty(t *testing.T) {

	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Test IsEmpty on an empty array
	if !arr.IsEmpty() {
		t.Errorf("Expected IsEmpty to return true, got false")
	}

	// Test IsEmpty on a non-empty array
	arr.InsertAt(0, 10)
	if arr.IsEmpty() {
		t.Errorf("Expected IsEmpty to return false, got true")
	}
}

func TestLen(t *testing.T) {

	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Test Len on an empty array
	if arr.Len() != 0 {
		t.Errorf("Expected Len to return 0, got %v", arr.Len())
	}

	// Test Len on a non-empty array
	arr.InsertAt(0, 10)
	if arr.Len() != 1 {
		t.Errorf("Expected Len to return 1, got %v", arr.Len())
	}
}

func TestToArray(t *testing.T) {

	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Insert some values
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(2, 30)

	// Test ToArray
	if len(arr.ToArray()) != 3 {
		t.Errorf("Expected ToArray to return an array of length 3, got %v", len(arr.ToArray()))
	}
	if arr.ToArray()[0] != 10 {
		t.Errorf("Expected value at index 0 to be 10, got %v", arr.ToArray()[0])
	}
	if arr.ToArray()[1] != 20 {
		t.Errorf("Expected value at index 1 to be 20, got %v", arr.ToArray()[1])
	}
	if arr.ToArray()[2] != 30 {
		t.Errorf("Expected value at index 2 to be 30, got %v", arr.ToArray()[2])
	}
}

func TestValues(t *testing.T) {

	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Insert some values
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(2, 30)

	// Test Values
	if len(arr.ToArray()) != 3 {
		t.Errorf("Expected Values to return an array of length 3, got %v", len(arr.ToArray()))
	}
	if arr.ToArray()[0] != 10 {
		t.Errorf("Expected value at index 0 to be 10, got %v", arr.ToArray()[0])
	}
	if arr.ToArray()[1] != 20 {
		t.Errorf("Expected value at index 1 to be 20, got %v", arr.ToArray()[1])
	}
	if arr.ToArray()[2] != 30 {
		t.Errorf("Expected value at index 2 to be 30, got %v", arr.ToArray()[2])
	}
}

func TestNewDynamicArray(t *testing.T) {

	// Create a new instance of the Array struct
	arr := NewDynamicArray()

	// Test NewDynamicArray
	if len(arr.ToArray()) != 0 {
		t.Errorf("Expected Values to return an array of length 0, got %v", len(arr.ToArray()))
	}
}
