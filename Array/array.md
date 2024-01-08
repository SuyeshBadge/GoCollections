## Array Package Documentation

### Introduction

The `array` package provides a versatile implementation of a dynamic and static array in Go. This package is useful for managing collections of elements and includes various methods for manipulation and retrieval.

### Array Structure

The central structure in this package is the `Array` type, which is defined as follows:

```go
type Array struct {
	values   []interface{}
	isStatic bool
}
```

- `values`: A slice holding the elements of the array.
- `isStatic`: A boolean indicating whether the array is static (with a fixed size) or dynamic.

### Creating Arrays

#### NewStaticArray

```go
func NewStaticArray(size int) *Array
```

Creates a new static array with the specified size.

- Parameters:
  - `size`: The size of the static array.

#### NewDynamicArray

```go
func NewDynamicArray() *Array
```

Creates a new dynamic array with a default size.

### Array Manipulation

#### Push

```go
func (a *Array) Push(value interface{}) error
```

Adds an element to the end of the array.

- Parameters:
  - `value`: The element to be added.

#### Pop

```go
func (a *Array) Pop() interface{}
```

Removes and returns the last element from the array.

#### ToArray

```go
func (a *Array) ToArray() []interface{}
```

Returns a copy of the array as a slice.

#### Len

```go
func (a *Array) Len() int
```

Returns the current length of the array.

#### IsEmpty

```go
func (a *Array) IsEmpty() bool
```

Returns `true` if the array is empty, `false` otherwise.

#### Contains

```go
func (a *Array) Contains(value interface{}) bool
```

Checks if the array contains the specified element.

#### IndexOf

```go
func (a *Array) IndexOf(value interface{}) int
```

Returns the index of the first occurrence of the specified element, or -1 if not found.

#### InsertAt

```go
func (a *Array) InsertAt(index int, value interface{}) error
```

Inserts an element at the specified index.

- Parameters:
  - `index`: The index at which to insert the element.
  - `value`: The element to be inserted.

#### RemoveAt

```go
func (a *Array) RemoveAt(index int) error
```

Removes the element at the specified index.

- Parameters:
  - `index`: The index of the element to be removed.

#### Clear

```go
func (a *Array) Clear()
```

Clears all elements from the array.

### Array Information

#### Capacity

```go
func (a *Array) Capacity() int
```

Returns the current capacity of the array.

#### Resize

```go
func (a *Array) Resize(newSize int)
```

Resizes the array to the specified new size.

- Parameters:
  - `newSize`: The new size of the array.

### Array Access

#### Get

```go
func (a *Array) Get(index int) (interface{}, error)
```

Returns the element at the specified index.

- Parameters:
  - `index`: The index of the element to be retrieved.

#### Set

```go
func (a *Array) Set(index int, value interface{}) error
```

Sets the value of the element at the specified index.

- Parameters:
  - `index`: The index of the element to be modified.
  - `value`: The new value for the element.

### Error Handling

All methods that may encounter errors return an `error` value, allowing for proper error handling in the calling code.

### Example Usage

```go
package main

import (
	"fmt"
	"array"
)

func main() {
	// Create a dynamic array
	arr := array.NewDynamicArray()

	// Push elements
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)

	// Print array length
	fmt.Println("Array Length:", arr.Len())

	// Print array elements
	fmt.Println("Array Elements:", arr.ToArray())
}
```

### Conclusion

This `array` package provides a flexible and efficient array implementation in Go, suitable for various use cases, including dynamic and static arrays. Users can leverage its methods to manipulate, access, and manage elements within the array.
