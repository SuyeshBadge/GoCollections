package linkedList

import (
	"testing"
)

func TestLinkedList_RemoveAt(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test removing an element at index 0
	removed := ll.RemoveAt(0)

	if removed != nil && removed.value != 10 {
		t.Errorf("Expected removed element to be 10, got %v", removed.value)
	}

	// Test removing an element at index 2
	removed = ll.RemoveAt(2)
	if removed.value != nil && removed.value != 40 {
		t.Errorf("Expected removed element to be 30, got %v", removed.value)
	}

	// Test removing an element at index 4
	removed = ll.RemoveAt(4)
	if removed != nil && removed.value != 50 {
		t.Errorf("Expected removed element to be 50, got %v", removed.value)
	}

	// Test removing an element at index out of range
	removed = ll.RemoveAt(10)
	if removed != nil {
		t.Errorf("Expected removed element to be nil, got %v", removed)
	}

	// Test removing an element from an empty list
	emptyList := NewLinkedList()
	removed = emptyList.RemoveAt(0)
	if removed != nil {
		t.Errorf("Expected removed element to be nil, got %v", removed)
	}
}

func TestLinkedList_Insert(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test inserting an element at index 0
	inserted := ll.Insert(0, 5)
	if inserted != true {
		t.Errorf("Expected inserted to be true, got %v", inserted)
	}

	// Test inserting an element at index 2
	inserted = ll.Insert(2, 25)
	if inserted != true {
		t.Errorf("Expected inserted to be true, got %v", inserted)
	}

	// Test inserting an element at index 6
	inserted = ll.Insert(6, 55)
	if inserted != true {
		t.Errorf("Expected inserted to be true, got %v", inserted)
	}

	// Test inserting an element at index out of range
	inserted = ll.Insert(10, 100)
	if inserted != false {
		t.Errorf("Expected inserted to be false, got %v", inserted)
	}

	// Test inserting an element at index -1
	inserted = ll.Insert(-1, 0)
	if inserted != false {
		t.Errorf("Expected inserted to be false, got %v", inserted)
	}

	// Test inserting an element into an empty list
	emptyList := NewLinkedList()
	inserted = emptyList.Insert(0, 10)
	if inserted != true {
		t.Errorf("Expected inserted to be true, got %v", inserted)
	}
}

func TestLinkedList_GetNode(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test getting an element at index 0
	node := ll.GetNode(0)
	if node == nil || node.value != 10 {
		t.Errorf("Expected node to be 10, got %v", node.value)
	}

	// Test getting an element at index 2
	node = ll.GetNode(2)
	if node == nil || node.value != 30 {
		t.Errorf("Expected node to be 30, got %v", node.value)
	}

	// Test getting an element at index 4
	node = ll.GetNode(4)
	if node == nil || node.value != 50 {
		t.Errorf("Expected node to be 50, got %v", node.value)
	}

	// Test getting an element at index out of range
	node = ll.GetNode(10)
	if node != nil {
		t.Errorf("Expected node to be nil, got %v", node)
	}

	// Test getting an element at index -1
	node = ll.GetNode(-1)
	if node != nil {
		t.Errorf("Expected node to be nil, got %v", node)
	}
}

func TestLinkedList_GetTail(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test getting the tail node
	node := ll.GetTail()
	if node == nil || node.value != 50 {
		t.Errorf("Expected node to be 50, got %v", node.value)
	}

	// Test getting the tail node from an empty list
	emptyList := NewLinkedList()
	node = emptyList.GetTail()
	if node != nil {
		t.Errorf("Expected node to be nil, got %v", node)
	}
}

func TestLinkedList_GetHead(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test getting the head node
	node := ll.GetHead()
	if node == nil || node.value != 10 {
		t.Errorf("Expected node to be 10, got %v", node.value)
	}

	// Test getting the head node from an empty list
	emptyList := NewLinkedList()
	node = emptyList.GetHead()
	if node != nil {
		t.Errorf("Expected node to be nil, got %v", node)
	}
}

func TestLinkedList_Add(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Test adding an element to an empty list
	added := ll.Add(10)
	if added != true {
		t.Errorf("Expected added to be true, got %v", added)
	}

	// Test adding an element to a non-empty list
	added = ll.Add(20)
	if added != true {
		t.Errorf("Expected added to be true, got %v", added)
	}
}

func TestLinkedList_Size(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test getting the size of the linked list
	size := ll.Size()
	if size != 5 {
		t.Errorf("Expected size to be 5, got %v", size)
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {

	// Create a new linked list
	ll := NewLinkedList()

	// Test checking if the linked list is empty
	empty := ll.IsEmpty()
	if empty != true {
		t.Errorf("Expected empty to be true, got %v", empty)
	}

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test checking if the linked list is empty
	empty = ll.IsEmpty()
	if empty != false {
		t.Errorf("Expected empty to be false, got %v", empty)
	}
}

func TestLinkedList_Clear(t *testing.T) {

	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Clear the linked list
	ll.Clear()

	// Test checking if the linked list is empty
	empty := ll.IsEmpty()
	if empty != true {
		t.Errorf("Expected empty to be true, got %v", empty)
	}
}

func TestLinkedList_Values(t *testing.T) {

	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Get the values from the linked list
	values := ll.Values()

	// Test getting the values from the linked list
	if len(values) != 5 {
		t.Errorf("Expected values length to be 5, got %v", len(values))
	}
}

func TestLinkedList_String(t *testing.T) {

	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Get the string representation of the linked list
	str := ll.String()

	// Test getting the string representation of the linked list
	if str != "[ 10 20 30 40 50 ]" {
		t.Errorf("Expected str to be [ 10 20 30 40 50 ], got %v", str)
	}
}

func TestLinkedList_Contains(t *testing.T) {

	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add("20")
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test checking if the linked list contains an element
	contains := ll.Contains("20")
	if contains != true {
		t.Errorf("Expected contains to be true, got %v", contains)
	}

	// Test checking if the linked list contains an element
	contains = ll.Contains(60)
	if contains != false {
		t.Errorf("Expected contains to be false, got %v", contains)
	}
}

func TestLinkedList_IndexOf(t *testing.T) {

	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add("20")
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test getting the index of an element
	index := ll.IndexOf("20")
	if index != 1 {
		t.Errorf("Expected index to be 1, got %v", index)
	}

	// Test getting the index of an element
	index = ll.IndexOf(60)
	if index != -1 {
		t.Errorf("Expected index to be -1, got %v", index)
	}
}

func TestLinkedList_LastIndexOf(t *testing.T) {

	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add("20")
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)
	ll.Add("20")

	// Test getting the last index of an element
	index := ll.LastIndexOf("20")
	if index != 5 {
		t.Errorf("Expected index to be 5, got %v", index)
	}

	// Test getting the last index of an element
	index = ll.LastIndexOf(60)
	if index != -1 {
		t.Errorf("Expected index to be -1, got %v", index)
	}
}

func TestLinkedList_Remove(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test removing an element from the linked list
	removed := ll.Remove(30)
	if removed != true {
		t.Errorf("Expected removed to be true, got %v", removed)
	}

	// Test removing the head element from the linked list
	removed = ll.Remove(10)
	if removed != true {
		t.Errorf("Expected removed to be true, got %v", removed)
	}

	// Test removing the tail element from the linked list
	removed = ll.Remove(50)
	if removed != true {
		t.Errorf("Expected removed to be true, got %v", removed)
	}

	// Test removing a non-existent element from the linked list
	removed = ll.Remove(100)
	if removed != false {
		t.Errorf("Expected removed to be false, got %v", removed)
	}

	// Test removing an element from an empty list
	emptyList := NewLinkedList()
	removed = emptyList.Remove(10)
	if removed != false {
		t.Errorf("Expected removed to be false, got %v", removed)
	}
}

func TestLinkedList_Equals(t *testing.T) {
	// Create a new linked list
	ll1 := NewLinkedList()
	ll2 := NewLinkedList()

	// Add some elements to the linked lists
	ll1.Add(10)
	ll1.Add(20)
	ll1.Add(30)

	ll2.Add(10)
	ll2.Add(20)
	ll2.Add(30)

	// Test when the linked lists are equal
	equal := ll1.Equals(ll2)
	if !equal {
		t.Errorf("Expected linked lists to be equal")
	}

	// Test when the linked lists are not equal
	ll2.Add(40)
	notEqual := ll1.Equals(ll2)
	if notEqual {
		t.Errorf("Expected linked lists to be not equal")
	}
}
func TestLinkedList_Copy(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Make a copy of the linked list
	copy := ll.Copy()

	// Test if the copy is equal to the original linked list
	if !ll.Equals(copy) {
		t.Errorf("Expected the copied linked list to be equal to the original linked list")
	}

	// Modify the original linked list
	ll.RemoveAt(0)
	ll.Insert(2, 35)

	// Test if the copy is still equal to the original linked list
	if ll.Equals(copy) {
		t.Errorf("Expected the copied linked list to be different from the modified original linked list")
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Reverse the linked list
	ll.Reverse()

	// Test the reversed linked list
	expected := []int{50, 40, 30, 20, 10}
	for i, val := range expected {
		node := ll.GetNode(i)
		if node == nil || node.value != val {
			t.Errorf("Expected node value at index %d to be %d, got %v", i, val, node.value)
		}
	}
}

func TestLinkedList_GetMiddle(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test getting the middle node
	middle := ll.GetMiddle()
	if middle == nil || middle.value != 30 {
		t.Errorf("Expected middle node to be 30, got %v", middle.value)
	}

	// Test getting the middle node from an empty list
	emptyList := NewLinkedList()
	middle = emptyList.GetMiddle()
	if middle != nil {
		t.Errorf("Expected middle node to be nil, got %v", middle)
	}

	// Test getting the middle node from a list with one element
	singleElementList := NewLinkedList()
	singleElementList.Add(10)
	middle = singleElementList.GetMiddle()
	if middle == nil || middle.value != 10 {
		t.Errorf("Expected middle node to be 10, got %v", middle.value)
	}

	// Test getting the middle node from a list with even number of elements
	evenElementsList := NewLinkedList()
	evenElementsList.Add(10)
	evenElementsList.Add(20)
	evenElementsList.Add(30)
	evenElementsList.Add(40)
	middle = evenElementsList.GetMiddle()
	if middle == nil || middle.value != 20 {
		t.Errorf("Expected middle node to be 20, got %v", middle.value)
	}
}

func TestLinkedList_GetNthFromEnd(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)

	// Test getting the Nth node from the end of the linked list
	// Test case 1: Get the 2nd node from the end (index 3)
	node := ll.GetNthFromEnd(2)
	if node == nil || node.value != 30 {
		t.Errorf("Expected node to be 30, got %v", node.value)
	}

	// Test case 2: Get the 4th node from the end (index 1)
	node = ll.GetNthFromEnd(4)
	if node == nil || node.value != 10 {
		t.Errorf("Expected node to be 10, got %v", node.value)
	}

	// Test case 3: Get the 6th node from the end (index out of range)
	node = ll.GetNthFromEnd(6)
	if node != nil {
		t.Errorf("Expected node to be nil, got %v", node)
	}

	// Test case 4: Get the -1st node from the end (index out of range)
	node = ll.GetNthFromEnd(-1)
	if node != nil {
		t.Errorf("Expected node to be nil, got %v", node)
	}

}
func TestLinkedList_RemoveDuplicates(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Add some elements to the linked list
	ll.RemoveDuplicates()
	ll.Add(10)
	ll.Add(20)
	ll.Add(20)
	ll.Add(30)
	ll.Add(30)
	ll.Add(30)
	ll.Add(40)
	ll.Add(50)
	ll.Add(50)
	ll.Add(50)
	ll.Add(50)

	// Remove duplicates from the linked list
	ll.RemoveDuplicates()

	// Test the size of the linked list after removing duplicates
	size := ll.Size()
	if size != 5 {
		t.Errorf("Expected size to be 5, got %v", size)
	}

	// Test the values in the linked list after removing duplicates
	expectedValues := []int{10, 20, 30, 40, 50}
	for i, val := range expectedValues {
		node := ll.GetNode(i)
		if node == nil || node.value != val {
			t.Errorf("Expected node at index %d to be %d, got %v", i, val, node.value)
		}
	}
}
