# Set

Set is data structure implementation in Go (Golang).

A set is a data structure that can store any number of unique values in any order.

## **Install and Usage**

```
go get github.com/kushalkhadka7/Set

```

```
import "github.com/kushalkhadka7/Set"

```

## **Examples**

Initialization

```
s:= set.New()
```

## Some Operations

Add

```
s.Add("test")
```

Remove

```
s.Remove("test")
```

Size

```
size:= s.Size()
```

Union

```
ns:= set.New()

ns.Add("test1")

union := s.Union(ns)
```

Intersection

```
ns:= set.New()

ns.Add("test1")

i := s.Intersection(ns)
```

## **Available Methods**

```
    Clear  # Empty current set.

	Size   # Size returns the current size of the set.

	IsEmpty  # IsEmpty checks the current set has elements in it or not.

	List  # List converts the set into slices.

	Union # Union returns a new set with all the unique values in current set and other set.

	Has  # Has checks either the given item is in the set or not.

	IsSubSet  # IsSubSet checks either the current set is a subset of the other given set.

	Add  # Add insert the given values in the set.

	Remove # Removes the given values from the set.

	Difference # Difference returns new set containing only the elements of current set.

	Intersection # Intersection returns a new set which contains items that only exist in the current set.
```
