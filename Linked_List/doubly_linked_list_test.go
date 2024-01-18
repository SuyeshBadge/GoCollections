package linked_list

import (
	"testing"
)

func TestDoublyLinkedListInsertAt(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test inserting at index 0 when the list is empty
	err := list.InsertAt(0, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}

	value, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %v", value)
	}

	// Test inserting at index 0 when the list is not empty
	err = list.InsertAt(0, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2, got %d", list.Size())
	}
	element0, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if element0 != 2 {
		t.Errorf("Expected element at index 0 to be 2, got %v", element0)
	}
	value, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 1 {
		t.Errorf("Expected element at index 1 to be 1, got %v", value)
	}

	// Test inserting at the end of the list
	err = list.InsertAt(2, 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 3 {
		t.Errorf("Expected list size to be 3, got %d", list.Size())
	}
	value, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 3 {
		t.Errorf("Expected element at index 2 to be 3, got %v", value)
	}

	err = list.InsertAt(2, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 4 {
		t.Errorf("Expected list size to be 4, got %d", list.Size())
	}
	value, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 4 {
		t.Errorf("Expected element at index 2 to be 4, got %v", value)
	}

	// Test inserting at an invalid index
	err = list.InsertAt(5, 4)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}

	if list.Size() != 4 {
		t.Errorf("Expected list size to be 3, got %d", list.Size())
	}
}

func TestDoublyLinkedListInsertAtInvalidIndex(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test inserting at a negative index
	err := list.InsertAt(-1, 1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}

	// Test inserting at an index greater than the list size
	err = list.InsertAt(1, 2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}
}

func TestDoublyLinkedListRemoveAt(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test removing from an empty list
	err := list.RemoveAt(0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}

	// Test removing from a list with one element
	list.InsertAt(0, 1)
	err = list.RemoveAt(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}

	// Test removing from a list with multiple elements
	list.InsertAt(0, 1)
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	err = list.RemoveAt(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2, got %d", list.Size())
	}
	value, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %v", value)
	}
	value, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 3 {
		t.Errorf("Expected element at index 1 to be 3, got %v", value)
	}

	// Test removing from the end of the list
	err = list.RemoveAt(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}
	value, err = list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %v", value)
	}

	// Test removing from the beginning of the list
	err = list.RemoveAt(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}
}

func TestDoublyLinkedListRemoveAtInvalidIndex(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test removing at a negative index
	err := list.RemoveAt(-1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}

	// Test removing at an index greater than the list size
	err = list.RemoveAt(1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}
}

func TestDoublyLinkedListGet(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test getting from an empty list
	_, err := list.Get(0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}

	// Test getting from a list with one element
	list.InsertAt(0, 1)
	value, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %v", value)
	}

	// Test getting from a list with multiple elements
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	value, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %v", value)
	}
}

func TestDoublyLinkedListGetInvalidIndex(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test getting at a negative index
	_, err := list.Get(-1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}

	// Test getting at an index greater than the list size
	_, err = list.Get(1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}
}

func TestDoublyLinkedListSize(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test size of an empty list
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}

	// Test size of a list with one element
	list.InsertAt(0, 1)
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}

	// Test size of a list with multiple elements
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	if list.Size() != 3 {
		t.Errorf("Expected list size to be 3, got %d", list.Size())
	}
}

func TestDoublyLinkedListIsEmpty(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test IsEmpty on an empty list
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty, got false")
	}

	// Test IsEmpty on a list with one element
	list.InsertAt(0, 1)
	if list.IsEmpty() {
		t.Errorf("Expected list to not be empty, got true")
	}

	// Test IsEmpty on a list with multiple elements
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	if list.IsEmpty() {
		t.Errorf("Expected list to not be empty, got true")
	}
}

func TestDoublyLinkedListClear(t *testing.T) {

	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test clearing an empty list
	list.Clear()
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}

	// Test clearing a list with one element
	list.InsertAt(0, 1)
	list.Clear()
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}

	// Test clearing a list with multiple elements
	list.InsertAt(0, 1)
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	list.Clear()
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}
}

func TestDoublyLinkedListContains(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test Contains on an empty list
	if list.Contains(1) {
		t.Errorf("Expected list to not contain 1, got true")
	}

	// Test Contains on a list with one element
	list.InsertAt(0, 1)
	if !list.Contains(1) {
		t.Errorf("Expected list to contain 1, got false")
	}
	if list.Contains(2) {
		t.Errorf("Expected list to not contain 2, got true")
	}

	// Test Contains on a list with multiple elements
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	if !list.Contains(2) {
		t.Errorf("Expected list to contain 2, got false")
	}
	if list.Contains(4) {
		t.Errorf("Expected list to not contain 4, got true")
	}
}

func TestDoublyLinkedListSearch(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test Search on an empty list
	if list.Search(1) != -1 {
		t.Errorf("Expected list to not contain 1, got true")
	}

	// Test Search on a list with one element
	list.InsertAt(0, 1)
	if list.Search(1) != 0 {
		t.Errorf("Expected list to contain 1, got false")
	}
	if list.Search(2) != -1 {
		t.Errorf("Expected list to not contain 2, got true")
	}

	// Test Search on a list with multiple elements
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	if list.Search(2) != 1 {
		t.Errorf("Expected list to contain 2, got false")
	}
	if list.Search(4) != -1 {
		t.Errorf("Expected list to not contain 4, got true")
	}
}

func TestDoublyLinkedListValues(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test Values on an empty list
	values := list.Values()
	if len(values) != 0 {
		t.Errorf("Expected values slice to be empty, got %v", values)
	}

	// Test Values on a list with one element
	list.InsertAt(0, 1)
	values = list.Values()
	if len(values) != 1 {
		t.Errorf("Expected values slice to have length 1, got %v", values)
	}
	if values[0] != 1 {
		t.Errorf("Expected values slice to contain 1, got %v", values)
	}

	// Test Values on a list with multiple elements
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	values = list.Values()
	if len(values) != 3 {
		t.Errorf("Expected values slice to have length 3, got %v", values)
	}
	if values[0] != 1 {
		t.Errorf("Expected values slice to contain 1, got %v", values)
	}
	if values[1] != 2 {
		t.Errorf("Expected values slice to contain 2, got %v", values)
	}
	if values[2] != 3 {
		t.Errorf("Expected values slice to contain 3, got %v", values)
	}
}

func TestDoublyLinkedListHead(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test Head on an empty list
	if list.Head() != nil {
		t.Errorf("Expected head to be nil, got %v", list.Head())
	}

	// Test Head on a list with one element
	list.InsertAt(0, 1)
	if list.Head().value != 1 {
		t.Errorf("Expected head to be 1, got %v", list.Head())
	}

	// Test Head on a list with multiple elements
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	if list.Head().value != 1 {
		t.Errorf("Expected head to be 1, got %v", list.Head())
	}
}

func TestDoublyLinkedListTail(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test Tail on an empty list
	if list.Tail() != nil {
		t.Errorf("Expected tail to be nil, got %v", list.Tail())
	}

	// Test Tail on a list with one element
	list.InsertAt(0, 1)
	if list.Tail().value != 1 {
		t.Errorf("Expected tail to be 1, got %v", list.Tail())
	}

	// Test Tail on a list with multiple elements
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	if list.Tail().value != 3 {
		t.Errorf("Expected tail to be 3, got %v", list.Tail())
	}
}

// add,remove,set
func TestDoublyLinkedListAdd(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test Add on an empty list
	list.Add(1)
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}

	// Test Add on a list with one element
	list.Add(2)
	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2, got %d", list.Size())
	}

	// Test Add on a list with multiple elements
	list.Add(3)
	if list.Size() != 3 {
		t.Errorf("Expected list size to be 3, got %d", list.Size())
	}
}

func TestDoublyLinkedListRemove(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test Remove on an empty list
	list.Remove(1)
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}

	// Test Remove on a list with one element
	list.Add(1)
	list.Remove(1)
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}

	// Test Remove on a list with multiple elements
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Remove(2)
	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2, got %d", list.Size())
	}
}

func TestDoublyLinkedListSet(t *testing.T) {
	// Create a new instance of the DoublyLinkedList struct
	list := NewDoublyLinkedList()

	// Test setting a value at index 0 when the list is empty will result in an error
	err := list.Set(0, 1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}
	list.Add(1)

	// Test setting a value at index 0 when the list is not empty
	err = list.Set(0, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}
	value, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 2 {
		t.Errorf("Expected element at index 0 to be 2, got %v", value)
	}

	// Test setting a value at the end of the list
	err = list.Set(1, 3)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}

	// Test setting a value at a valid index
	err = list.Set(0, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}
	value, err = list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if value != 4 {
		t.Errorf("Expected element at index 0 to be 4, got %v", value)
	}

	// Test setting a value at an invalid index
	err = list.Set(2, 5)
	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Error() != "index out of range" {
		t.Errorf("Expected error to be 'index out of range', got %v", err)
	}

	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}
}
