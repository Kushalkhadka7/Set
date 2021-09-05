package set

type Set struct{}

type Setter interface {
	Add() (string, error)
}

func NewSet() Setter {
	return &Set{}
}

func (s *Set) Add() (string, error) {
	return "hello world", nil
}
