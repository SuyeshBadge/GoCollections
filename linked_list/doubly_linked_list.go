package linked_list

import "errors"

// DoublyLinkedListNode represents a node in a doubly linked list.
type DoublyLinkedListNode struct {
	value interface{}
	next  *DoublyLinkedListNode
	prev  *DoublyLinkedListNode
}

// DoublyLinkedList represents a doubly linked list data structure.
type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
	size int
}

// NewDoublyLinkedList creates and returns a new instance of DoublyLinkedList.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// InsertAt inserts a new node at the specified index in the doubly linked list.
func (l *DoublyLinkedList) InsertAt(index int, value interface{}) error {
	if index < 0 || index > l.Size() {
		return errors.New("index out of range")
	}

	newNode := &DoublyLinkedListNode{value: value}

	if l.Size() == 0 {
		l.head = newNode
		l.tail = newNode
	} else if index == 0 {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	} else if index == l.Size() {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	} else {
		currentNode := l.head
		for i := 0; i < index-1; i++ {
			currentNode = currentNode.next
		}
		newNode.next = currentNode.next
		newNode.prev = currentNode
		currentNode.next.prev = newNode
		currentNode.next = newNode
	}

	l.size++

	return nil
}

// RemoveAt removes a node at the specified index in the doubly linked list.
func (l *DoublyLinkedList) RemoveAt(index int) error {
	if index < 0 || index >= l.Size() {
		return errors.New("index out of range")
	}

	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
	} else if index == 0 {
		l.head = l.head.next
		l.head.prev = nil
	} else if index == l.Size()-1 {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		currentNode := l.head
		for i := 0; i < index-1; i++ {
			currentNode = currentNode.next
		}
		currentNode.next = currentNode.next.next
		currentNode.next.prev = currentNode
	}

	l.size--

	return nil
}

// Size returns the number of elements in the doubly linked list.
func (l *DoublyLinkedList) Size() int {
	return l.size
}

// IsEmpty returns true if the doubly linked list is empty, otherwise returns false.
func (l *DoublyLinkedList) IsEmpty() bool {
	return l.Size() == 0
}

// Head returns the head node of the doubly linked list.
func (l *DoublyLinkedList) Head() *DoublyLinkedListNode {
	return l.head
}

// Tail returns the tail node of the doubly linked list.
func (l *DoublyLinkedList) Tail() *DoublyLinkedListNode {
	return l.tail
}

// Search returns the index of the first occurrence of the specified value in the doubly linked list, otherwise returns -1.
func (l *DoublyLinkedList) Search(value interface{}) int {
	currentNode := l.head
	for i := 0; i < l.Size(); i++ {
		if currentNode.value == value {
			return i
		}
		currentNode = currentNode.next
	}
	return -1
}

// Values returns a slice of all values in the doubly linked list.
func (l *DoublyLinkedList) Values() []interface{} {
	values := make([]interface{}, l.Size())
	currentNode := l.head
	for i := 0; i < l.Size(); i++ {
		values[i] = currentNode.value
		currentNode = currentNode.next
	}
	return values
}

// Add adds a new node with the specified value to the end of the doubly linked list.
// The value parameter represents the value to be added to the list.
func (l *DoublyLinkedList) Add(value interface{}) {
	l.InsertAt(l.Size(), value)
}

// Remove removes the first occurrence of the specified value from the doubly linked list.
// If the value is not found, it returns an error.
func (l *DoublyLinkedList) Remove(value interface{}) error {
	index := l.Search(value)
	if index == -1 {
		return errors.New("value not found")
	}
	return l.RemoveAt(index)
}

// Contains checks if the doubly linked list contains the specified value.
// It returns true if the value is found, otherwise it returns false.
func (l *DoublyLinkedList) Contains(value interface{}) bool {
	return l.Search(value) != -1
}

// Clear removes all elements from the doubly linked list.
// It sets the head and tail pointers to nil and resets the size to 0.
func (l *DoublyLinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Get returns the value at the specified index in the doubly linked list.
// If the index is out of range, it returns an error.
func (l *DoublyLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.Size() {
		return nil, errors.New("index out of range")
	}
	currentNode := l.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.value, nil
}

// Set sets the value at the specified index in the doubly linked list.
// If the index is out of range, it returns an error.
func (l *DoublyLinkedList) Set(index int, value interface{}) error {
	if index < 0 || index >= l.Size() {
		return errors.New("index out of range")
	}
	currentNode := l.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	currentNode.value = value
	return nil
}
