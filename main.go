package main

import (
	"fmt"

	"github.com/Kushalkhadka7/Set/set"
)

func main() {
	s := set.New()
	s.Add("hello")
	s.Add("world")

	fmt.Println(s)
	fmt.Println(s.Has("world", "hello"))
	fmt.Println(s)
	fmt.Println(s.Size())
	s.Clear()
	fmt.Println(s.Size())
}
