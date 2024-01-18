# Doubly Linked List

## Package Name: linked_list

### Overview

The `linked_list` package provides a robust implementation of a doubly linked list data structure in Go. A doubly linked list is a collection of elements, each containing a value and pointers to the next and previous elements in the sequence. This package supports various operations such as insertion, removal, search, and traversal of the doubly linked list.

### Types

#### `DoublyLinkedListNode`

```go
type DoublyLinkedListNode struct {
	value interface{}
	next  *DoublyLinkedListNode
	prev  *DoublyLinkedListNode
}
```

Represents a node in the doubly linked list. Each node contains a value, a pointer to the next node, and a pointer to the previous node.

#### `DoublyLinkedList`

```go
type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
	size int
}
```

Represents the doubly linked list itself, containing pointers to the head and tail nodes, as well as the size (number of elements) in the list.

### Functions

#### `NewDoublyLinkedList`

```go
func NewDoublyLinkedList() *DoublyLinkedList
```

Creates and returns a new instance of `DoublyLinkedList`.

#### `InsertAt`

```go
func (l *DoublyLinkedList) InsertAt(index int, value interface{}) error
```

Inserts a new node with the specified value at the given index in the doubly linked list.

#### `RemoveAt`

```go
func (l *DoublyLinkedList) RemoveAt(index int) error
```

Removes the node at the specified index from the doubly linked list.

#### `Size`

```go
func (l *DoublyLinkedList) Size() int
```

Returns the number of elements in the doubly linked list.

#### `IsEmpty`

```go
func (l *DoublyLinkedList) IsEmpty() bool
```

Returns `true` if the doubly linked list is empty; otherwise, returns `false`.

#### `Head`

```go
func (l *DoublyLinkedList) Head() *DoublyLinkedListNode
```

Returns the head node of the doubly linked list.

#### `Tail`

```go
func (l *DoublyLinkedList) Tail() *DoublyLinkedListNode
```

Returns the tail node of the doubly linked list.

#### `Search`

```go
func (l *DoublyLinkedList) Search(value interface{}) int
```

Returns the index of the first occurrence of the specified value in the doubly linked list. Returns -1 if the value is not found.

#### `Values`

```go
func (l *DoublyLinkedList) Values() []interface{}
```

Returns a slice containing all values in the doubly linked list.

#### `Add`

```go
func (l *DoublyLinkedList) Add(value interface{})
```

Adds a new node with the specified value to the end of the doubly linked list.

#### `Remove`

```go
func (l *DoublyLinkedList) Remove(value interface{}) error
```

Removes the first occurrence of the specified value from the doubly linked list. Returns an error if the value is not found.

#### `Contains`

```go
func (l *DoublyLinkedList) Contains(value interface{}) bool
```

Checks if the doubly linked list contains the specified value. Returns `true` if the value is found; otherwise, returns `false`.

#### `Clear`

```go
func (l *DoublyLinkedList) Clear()
```

Removes all elements from the doubly linked list, setting the head and tail pointers to `nil` and resetting the size to 0.

#### `Get`

```go
func (l *DoublyLinkedList) Get(index int) (interface{}, error)
```

Returns the value at the specified index in the doubly linked list. Returns an error if the index is out of range.

#### `Set`

```go
func (l *DoublyLinkedList) Set(index int, value interface{}) error
```

Sets the value at the specified index in the doubly linked list. Returns an error if the index is out of range.

### Example Usage

```go
package main

import (
	"fmt"
	"github.com/your-username/linked_list"
)

func main() {
	// Create a new doubly linked list
	list := linked_list.NewDoublyLinkedList()

	// Insert values at specific indices
	list.InsertAt(0, 10)
	list.InsertAt(1, 20)
	list.InsertAt(1, 15)

	// Print the values in the list
	fmt.Println("List values:", list.Values())

	// Remove a value at a specific index
	list.RemoveAt(1)

	// Print the updated list
	fmt.Println("Updated list values:", list.Values())
}
```

### License

This package is released under the [MIT License](https://opensource.org/licenses/MIT). See the LICENSE file for more details.