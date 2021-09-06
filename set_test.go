package set

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {

	testCases := []struct {
		name   string
		values []interface{}
		result Setter
	}{
		{
			name:   "Addig string types.",
			values: []interface{}{"Geeks", "hello", "world"},
			result: New().Add("Geeks", "hello", "world"),
		},
		{
			name:   "Addig mixed types.",
			values: []interface{}{"data", 1, nil},
			result: New().Add("data", 1, nil),
		},
		{
			name:   "Should have unique values.",
			values: []interface{}{"data", 1, nil, 2},
			result: New().Add("data", 1, nil, 2),
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.values {
				ok := tc.result.Has(v)
				if !ok {
					t.Errorf("%s", "Failed")
				}

			}

		})

	}
}

func TestRemove(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add("test")

	s.Remove("test")

	has := s.Has("test")

	if has {
		t.Error("Should not have test.")
	}

	if s.Size() != 2 {
		t.Error("After removing one element size should be 2.")
	}

	s.Remove(1)

	has = s.Has(1)

	if has {
		t.Error("Should not have test.")
	}

	if s.Size() != 1 {
		t.Error("After removing one element size should be 1.")
	}

}

func TestSize(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add("test")

	if s.Size() != 3 {
		t.Error("Size should be 3")
	}

	s.Remove(1)

	if s.Size() != 2 {
		t.Error("Size should be 2")
	}

	s.Remove(2)

	if s.Size() != 1 {
		t.Error("Size should be 1")
	}
}

func TestClear(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add("test")

	if s.Size() != 3 {
		t.Error("Size should be 3")
	}

	s.Clear()

	if s.Size() != 0 {
		t.Error("Size should be 0")
	}
}

func TestHas(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add("test")

	has := s.Has("Hello")
	require.Equal(t, has, false)

	has = s.Has("Test")
	require.Equal(t, has, false)

	has = s.Has("test")
	require.Equal(t, has, true)
}

func TestIsEmpty(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add("test")

	e := s.IsEmpty()
	require.Equal(t, e, false)

	s.Clear()

	e = s.IsEmpty()
	require.Equal(t, e, true)
}

func TestList(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add("test")

	list := []interface{}{1, 2, "test"}

	l := s.List()

	if len(list) != len(l) {
		t.Error("Not same size")
	}

	exists := make(map[interface{}]bool)
	for _, value := range list {
		exists[value] = true
	}

	for _, value := range l {
		if !exists[value] {
			t.Error("value are not same")
		}
	}
}

func TestInterection(t *testing.T) {
	a := New()
	a.Add(1)
	a.Add(2)
	a.Add("test")

	b := New()
	b.Add(1)
	a.Add(2)

	c := a.Intersection(b)

	if c.Size() == 2 {
		log.Println("Size not equal.")
	} else {
		t.Errorf("%s", "Size not equal.")
	}

}

func TestIsSubSet(t *testing.T) {
	a := New()
	a.Add(1)
	a.Add(2)
	a.Add("test")

	b := New()
	b.Add(1)
	b.Add(2)

	d := New()
	d.Add(1)
	d.Add(21)

	c := b.IsSubSet(a)
	e := d.IsSubSet(a)

	require.Equal(t, c, true)
	require.Equal(t, e, false)
}

func TestIsSuperSeet(t *testing.T) {
	a := New()
	a.Add(1)
	a.Add(2)
	a.Add("test")

	b := New()
	b.Add(1)
	b.Add(2)

	d := New()
	d.Add(1)
	d.Add(21)

	c := a.IsSubSet(b)
	e := a.IsSubSet(d)

	require.Equal(t, c, true)
	require.Equal(t, e, false)
}

func TestDifference(t *testing.T) {
	a := New()
	a.Add(1)
	a.Add(2)
	a.Add("test")

	b := New()
	b.Add(1)
	b.Add(2)

	c := b.Difference(a)

	if c.Has("test") {
		t.Error("Should not have test")
	}

	if !c.Has(1) {
		t.Error("Should have 1")
	}

}
