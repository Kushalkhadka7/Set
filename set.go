package set

import (
	"fmt"
	"sync"
)

// Set is a data structure that is used to store any number of unique values in any order.
type Set struct {
	data map[interface{}]interface{}
	sync.RWMutex
}

// Setter is an interface exposing all the methods that can be operated on Set.
type Setter interface {
	Clear()
	Size() int
	IsEmpty() bool
	List() []interface{}
	Union(otherSet Setter) Setter
	Has(item ...interface{}) bool
	IsSubSet(otherSet Setter) bool
	Add(items ...interface{}) Setter
	Remove(items ...interface{}) error
	Difference(otherSet Setter) Setter
	Intersection(otherSet Setter) Setter
}

// New creates and initalizes a new Set exposing the methods associated with it.
func New() Setter {
	return &Set{
		data: make(map[interface{}]interface{}),
	}
}

// Add insert the given values in the set.
// If passed nothing it will fail silently returns.
// If existing or repeated values are passed, then only the unique values are added.
func (s *Set) Add(items ...interface{}) Setter {
	if len(items) == 0 {
		return s
	}

	for _, i := range items {
		ok := s.Has(i)
		if ok {
			break
		}

		s.data[i] = true
	}

	return s
}

// Removes the given values from the set.
func (s *Set) Remove(items ...interface{}) error {
	if len(items) == 0 {
		return fmt.Errorf("%s", "No items provided to add.")
	}

	s.Lock()
	defer s.Unlock()

	for _, i := range items {
		_, ok := s.data[i]
		if !ok {
			panic(fmt.Errorf("not implemented"))
		}

		delete(s.data, i)
	}

	return nil
}

// Size returns the current size of the set.
// Counts the number of element in the set and returns the result.
func (s *Set) Size() int {
	s.Lock()
	defer s.Unlock()

	return len(s.data)
}

// Clear empty the set.
// Deletes all the values stored in the set.
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.data = make(map[interface{}]interface{})
}

// Has checks either the given item is in the set or not.
// Returns true if exists otherwise false.
func (s *Set) Has(items ...interface{}) bool {
	s.Lock()
	defer s.Unlock()

	ok := true

	for _, i := range items {
		_, ok = s.data[i]

		if !ok {
			break
		}

	}

	return ok
}

// IsEmpty checks the current set has elements in it or not.
func (s *Set) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()
	v := s.Size()

	return v == 0
}

// List converts the set into slices.
func (s *Set) List() []interface{} {
	s.Lock()
	defer s.Unlock()

	list := make([]interface{}, 0, len(s.data))

	for i := range s.data {
		list = append(list, i)
	}

	return list
}

// Intersection returns a new set which contains items that only exist in the current set.
func (s *Set) Intersection(otherSet Setter) Setter {
	newSet := New()

	for v := range s.data {
		val := otherSet.Has(v)

		if val {
			newSet.Add(v)
		}

		continue
	}

	return newSet
}

// Union returns a new set with all unique values in current set and other set.
func (s *Set) Union(otherSet Setter) Setter {
	newSet := New()

	for _, v := range s.List() {

		for _, vl := range otherSet.List() {
			if vl == v {
				continue
			}

			newSet.Add(vl)
		}

		if newSet.Has(v) {
			continue
		}

		newSet.Add(v)
	}

	return newSet
}

// IsSubSet checks either the current set is a subset of the other given set.
// Set s is a subset of otherSet if all elements of the set s are elements of the otherSet.
func (s *Set) IsSubSet(otherSet Setter) bool {
	var isSubset = false

	for _, v := range s.List() {
		has := otherSet.Has(v)
		fmt.Printf("value:%v\n", has)

		if has {
			isSubset = true

			continue
		}

		isSubset = false

		break
	}

	return isSubset
}

// IsSuperSet checks either the current set is a superset of the other given set.
// Set s is a superset of otherSet if all elements of the otherSet are elements of the set s.
func (s *Set) IsSuperSet(otherSet Setter) bool {
	var isSuperSet = false

	for _, v := range otherSet.List() {
		has := s.Has(v)

		if has {
			isSuperSet = true

			continue
		}

		isSuperSet = false

		break
	}

	return isSuperSet
}

// Difference returns new set containing only the elements of current set.
func (s *Set) Difference(otherSet Setter) Setter {
	newSet := New()

	for _, v := range s.List() {
		if otherSet.Has(v) {
			continue
		}

		newSet.Add(v)
	}

	return newSet
}
