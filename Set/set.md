# Set Package in Go

This package provides a Set data structure implementation in Go (Golang). A Set is a collection of distinct objects, called elements or members. This package provides various methods to manipulate sets and perform common set operations.

## Usage

First, import the package in your Go file:

```go
import "set"
```

### Creating a Set

You can create a new set using the `NewSet` function:

```go
s := set.NewSet()
```

You can also create a new set from a slice using the `NewSetFromSlice` function:

```go
s := set.NewSetFromSlice([]interface{}{1, 2, 3})
```

### Adding and Removing Elements

You can add an element to the set using the `Add` method:

```go
s.Add(4)
```

You can remove an element from the set using the `Remove` method:

```go
s.Remove(1)
```

### Checking if an Element Exists

You can check if an element exists in the set using the `Contains` method:

```go
exists := s.Contains(2)  // returns true if 2 is in the set
```

### Set Size

You can get the number of elements in the set using the `Len` method:

```go
size := s.Len()
```

### Clearing a Set

You can remove all elements from the set using the `Clear` method:

```go
s.Clear()
```

### Set Operations

The package provides functions to perform common set operations:

- Union: `Union(s1, s2 *Set) *Set`
- Intersection: `Intersection(s1, s2 *Set) *Set`
- Difference: `Difference(s1, s2 *Set) *Set`
- Symmetric Difference: `SymmetricDifference(s1, s2 *Set) *Set`
- Subset check: `IsSubset(s1, s2 *Set) bool`
- Superset check: `IsSuperset(s1, s2 *Set) bool`
- Disjoint check: `IsDisjoint(s1, s2 *Set) bool`

### Other Functions

The package also provides other useful functions:

- `Clone(s *Set) *Set`: Returns a new set that is a copy of the given set.
- `Equal(s1, s2 *Set) bool`: Checks if two sets are equal.
- `String() string`: Returns a string representation of the set.
- `IsEmpty() bool`: Checks if the set is empty.
- `IsProperSubset(s1, s2 *Set) bool`: Checks if s1 is a proper subset of s2.
- `IsProperSuperset(s1, s2 *Set) bool`: Checks if s1 is a proper superset of s2.
- `PowerSet(s *Set) []*Set`: Returns the power set of the given set.
- `CartesianProduct(s1, s2 *Set) []*Set`: Returns the Cartesian product of two sets.

## Note

The order of the elements in the set is not guaranteed. The `ToSlice` method returns a slice containing all the elements in the set, but the order of the elements in the slice is not guaranteed.

## License

This package is released under the MIT License.
