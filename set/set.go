package set

import (
	"fmt"
	"sync"
)

type Set struct {
	data map[interface{}]interface{}
	sync.RWMutex
}

type Setter interface {
	Has(item ...interface{}) bool
	Clear()
	Size() int
	Add(items ...interface{}) error
	Remove(items ...interface{}) error
}

func New() Setter {
	return &Set{
		data: make(map[interface{}]interface{}),
	}
}

func (s *Set) Add(items ...interface{}) error {
	if len(items) == 0 {
		return fmt.Errorf("%s", "No items provided to add.")
	}

	s.Lock()

	for _, i := range items {
		ok := s.Has(i)
		if ok {
			break
		}

		s.data[i] = true
	}

	s.Unlock()

	return nil
}

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

func (s *Set) Size() int {
	s.Lock()
	defer s.Unlock()

	return len(s.data)
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.data = make(map[interface{}]interface{})
}

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
