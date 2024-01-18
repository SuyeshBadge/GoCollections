# LinkedList Go Package

## Introduction

This is a simple implementation of a linked list in Go. A linked list is a data structure that consists of nodes where each node contains a value and a reference (link) to the next node in the sequence. The last node typically points to null, indicating the end of the list.

## Features

- **Addition and Removal**: Easily add and remove nodes with specified values.
- **Containment Check**: Check if the linked list contains a specific value.
- **Size and Empty Check**: Retrieve the number of nodes in the list and check if it's empty.
- **Clear Operation**: Remove all nodes from the linked list.
- **Conversion to Slice**: Obtain a slice containing the values in the linked list.
- **String Representation**: Get a string representation of the linked list.
- **Head and Tail Access**: Access the head and tail nodes of the linked list.
- **Node Access by Index**: Get the node at a specified index.
- **Insertion and Removal at Index**: Insert a new node or remove a node at a specified index.
- **Equality Check**: Compare two linked lists for equality.
- **Copy Operation**: Create a copy of the linked list.
- **Reverse Operation**: Reverse the order of nodes in the linked list.
- **Middle and Nth-from-End Node Access**: Get the middle node or the nth node from the end.
- **Duplicate Removal**: Remove duplicate nodes from the linked list.
- **Index Retrieval**: Get the index of the first and last occurrence of a specified value.

## Usage

### Creating a New LinkedList

```go
list := NewLinkedList()
```

### Adding and Removing Nodes

```go
list.Add(42)
list.Remove(42)
```

### Containment Check

```go
contains := list.Contains(42)
```

### Size and Empty Check

```go
size := list.Size()
isEmpty := list.IsEmpty()
```

### Clear Operation

```go
list.Clear()
```

### Conversion to Slice

```go
values := list.Values()
```

### String Representation

```go
str := list.String()
```

### Head and Tail Access

```go
head := list.GetHead()
tail := list.GetTail()
```

### Node Access by Index

```go
node := list.GetNode(2)
```

### Insertion and Removal at Index

```go
list.Insert(1, 42)
removedNode := list.RemoveAt(2)
```

### Equality Check

```go
list1 := NewLinkedList()
list2 := NewLinkedList()
isEqual := list1.Equals(list2)
```

### Copy Operation

```go
copy := list.Copy()
```

### Reverse Operation

```go
list.Reverse()
```

### Middle and Nth-from-End Node Access

```go
middleNode := list.GetMiddle()
nthNodeFromEnd := list.GetNthFromEnd(2)
```

### Duplicate Removal

```go
list.RemoveDuplicates()
```

### Index Retrieval

```go
indexOfValue := list.IndexOf(42)
lastIndexOfValue := list.LastIndexOf(42)
```

## Conclusion

This LinkedList package provides a flexible and convenient way to work with linked lists in Go, offering a variety of operations for list manipulation. It can be used in various scenarios where dynamic data storage with constant-time insertion and deletion at the beginning or end is required. Feel free to explore and use the functions provided to suit your specific use case.