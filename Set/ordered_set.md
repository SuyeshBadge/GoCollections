# OrderedSet Go Package

## Introduction

This is a simple implementation of an Ordered Set in Go. An Ordered Set is a data structure that maintains a collection of unique elements in the order they were added. It extends the functionality of a regular set by preserving the insertion order of elements.

## Features

- **Addition and Removal**: Easily add and remove elements from the set.
- **Access by Index**: Retrieve elements by their index in the set.
- **Set Operations**: Perform set operations like Union, Intersection, Difference, and more.
- **Power Set and Cartesian Product**: Calculate the power set and Cartesian product of the set.
- **Subset and Superset Checking**: Check if a set is a subset, superset, or equal to another set.
- **Disjoint Union**: Combine two sets into a disjoint union.
- **Pop Operation**: Remove and return an arbitrary item from the set.
- **Copy and Clone**: Create a copy or clone of the set.
- **String Representation**: Obtain a string representation of the set.

## Usage

### Importing the Package

```go
import ordered_set "goCollections/set"
```

### Creating a New Ordered Set

```go
set := ordered_set.NewOrderedSet()
```

### Adding and Removing Elements

```go
set.Add("apple")
set.Add("banana")
set.Remove("apple")
```

### Set Operations

```go
set1 := NewOrderedSet()
set1.Add(1)
set1.Add(2)

set2 := NewOrderedSet()
set2.Add(2)
set2.Add(3)

union := set1.Union(set2)
intersection := set1.Intersection(set2)
difference := set1.Difference(set2)
symmetricDifference := set1.SymmetricDifference(set2)
```

### Accessing Elements

```go
element := set.Get(0)
```

### Checking Subset, Superset, and Equality

```go
isSubset := set1.IsSubset(set2)
isSuperset := set1.IsSuperset(set2)
isEqual := set1.IsEqual(set2)
```

### Power Set and Cartesian Product

```go
powerSet := set.PowerSet()
cartesianProduct := set1.CartesianProduct(set2)
```

### Disjoint Union

```go
disjointUnion := set1.DisjointUnion(set2)
```

### Pop Operation

```go
item := set.Pop()
```

### Copy and Clone

```go
copy := set.Copy()
clone := set.Clone()
```

### Other Operations

```go
isEmpty := set.IsEmpty()
size := set.Len()
```

## Conclusion

This Ordered Set package provides a convenient and efficient way to work with ordered sets in Go, offering a variety of operations for set manipulation. It can be used in scenarios where maintaining element order is crucial. Feel free to explore and use the functions provided to suit your specific use case.
