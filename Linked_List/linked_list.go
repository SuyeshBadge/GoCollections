package linkedList

import "fmt"

// Node represents a node in a linked list.
type Node struct {
	value interface{} // The value stored in the node.
	next  *Node       // Pointer to the next node in the list.
}

// LinkedList represents a linked list data structure.
type LinkedList struct {
	head *Node // Pointer to the first node in the linked list
	size int   // Number of nodes in the linked list
}

// NewLinkedList creates and returns a new instance of LinkedList.
func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

// Add adds a new node with the specified value to the linked list.
// If the linked list is empty, the new node becomes the head of the list.
// Otherwise, the new node is inserted at the last of the list.
// The size of the linked list is incremented after adding the new node.
func (l *LinkedList) Add(value interface{}) bool {
	if l.head == nil {
		l.head = &Node{value, nil}
	} else {
		n := l.head
		for n.next != nil {
			n = n.next
		}
		n.next = &Node{value, nil}
	}
	l.size++
	return true
}

// Remove removes the first occurrence of the specified value from the linked list.
// If the value is found, it is removed and the size of the linked list is decremented.
// If the value is not found or the linked list is empty, no changes are made.
func (l *LinkedList) Remove(value interface{}) bool {
	if l.head == nil {
		return false
	}

	if l.head.value == value {
		l.head = l.head.next
		l.size--
		return true
	}

	prev := l.head
	for prev.next != nil {
		if prev.next.value == value {
			prev.next = prev.next.next
			l.size--
			return true
		}
		prev = prev.next
	}
	return false
}

// Contains checks if the linked list contains a specific value.
// It iterates through the linked list starting from the head node,
// comparing each node's value with the given value.
// If a node with the same value is found, it returns true.
// If the end of the linked list is reached without finding a matching value, it returns false.
func (l *LinkedList) Contains(value interface{}) bool {
	for n := l.head; n != nil; n = n.next {
		if n.value == value {
			return true
		}
	}
	return false
}

// Size returns the number of nodes in the linked list.
func (l *LinkedList) Size() int {
	return l.size
}

// IsEmpty returns true if the linked list is empty, false otherwise.
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// Clear removes all nodes from the linked list.
func (l *LinkedList) Clear() {
	l.head = nil
	l.size = 0
}

// Values returns a slice containing the values in the linked list.
func (l *LinkedList) Values() []interface{} {
	values := make([]interface{}, l.size)
	i := 0
	for n := l.head; n != nil; n = n.next {
		values[i] = n.value
		i++
	}
	return values
}

// String returns a string representation of the linked list.
func (l *LinkedList) String() string {
	str := "["
	for n := l.head; n != nil; n = n.next {
		str += fmt.Sprintf(" %v", n.value)
	}
	str += " ]"
	return str
}

// GetHead returns the head node of the linked list.
func (l *LinkedList) GetHead() *Node {
	return l.head
}

// GetTail returns the tail node of the linked list.
func (l *LinkedList) GetTail() *Node {
	if l.head == nil {
		return nil
	}

	n := l.head
	for n.next != nil {
		n = n.next
	}
	return n
}

// GetNode returns the node at the specified index.
// If the index is out of bounds, it returns nil.
func (l *LinkedList) GetNode(index int) *Node {
	if index < 0 || index >= l.size {
		return nil
	}

	n := l.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n
}

// Insert inserts a new node with the specified value at the specified index.
// If the index is out of bounds, it returns false.
// Otherwise, it returns true.
func (l *LinkedList) Insert(index int, value interface{}) bool {
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		l.Add(value)
		return true
	}

	n := l.head
	for i := 0; i < index-1; i++ {
		n = n.next
	}
	n.next = &Node{value, n.next}
	l.size++
	return true
}

// RemoveAt removes the node at the specified index.
// If the index is out of bounds, it returns nil.
// Otherwise, it returns the removed node.
func (l *LinkedList) RemoveAt(index int) *Node {
	if index < 0 || index >= l.size {
		return nil
	}

	if index == 0 {
		n := l.head
		l.head = l.head.next
		l.size--
		return n
	}

	n := l.head
	for i := 0; i < index-1; i++ {
		n = n.next
	}
	removed := n.next
	n.next = n.next.next
	l.size--
	return removed
}

// Equals compares the linked list with another linked list.
// It returns true if both linked lists have the same size and contain the same values in the same order.
// Otherwise, it returns false.
func (l *LinkedList) Equals(other *LinkedList) bool {
	if l.size != other.size {
		return false
	}

	n1 := l.head
	n2 := other.head
	for n1 != nil {
		if n1.value != n2.value {
			return false
		}
		n1 = n1.next
		n2 = n2.next
	}
	return true
}

// Copy returns a copy of the linked list.
// It creates a new linked list and copies the values from the original linked list to the new linked list.
// Returns a pointer to the new linked list.
func (l *LinkedList) Copy() *LinkedList {
	copy := NewLinkedList()
	for n := l.head; n != nil; n = n.next {
		copy.Add(n.value)
	}
	return copy
}

// Reverse reverses the linked list.
// It iterates through the linked list and swaps the next pointers of each node.
func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

// GetMiddle returns the middle node of the linked list.
// If the linked list has an even number of nodes, it returns the first middle node.
// If the linked list is empty, it returns nil.
func (l *LinkedList) GetMiddle() *Node {
	if l.head == nil {
		return nil
	}

	slow := l.head
	fast := l.head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// GetNthFromEnd returns the nth node from the end of the linked list.
// If the index is out of bounds, it returns nil.
func (l *LinkedList) GetNthFromEnd(index int) *Node {
	if index < 0 || index >= l.size {
		return nil
	}

	n := l.head
	for i := 0; i < l.size-index-1; i++ {
		n = n.next
	}
	return n
}

// RemoveDuplicates removes duplicate nodes from the linked list.
// It iterates through the linked list and removes nodes with duplicate values.
// If the linked list is empty, it returns nil.
func (l *LinkedList) RemoveDuplicates() {
	if l.head == nil {
		return
	}

	n := l.head
	for n != nil && n.next != nil {
		if n.value == n.next.value {
			n.next = n.next.next
			l.size--
		} else {
			n = n.next
		}
	}
}

// IndexOf returns the index of the first occurrence of the specified value in the linked list.
// If the value is found, the index is returned. If the value is not found, -1 is returned.
func (l *LinkedList) IndexOf(value interface{}) int {
	index := 0
	for n := l.head; n != nil; n = n.next {
		if n.value == value {
			return index
		}
		index++
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of the specified value in the linked list.
// If the value is found, the index is returned. If the value is not found, -1 is returned.
func (l *LinkedList) LastIndexOf(value interface{}) int {
	index := -1
	i := 0
	for n := l.head; n != nil; n = n.next {
		if n.value == value {
			index = i
		}
		i++
	}
	return index
}
