package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
func Print[T string | int](a []T) {
	for _, v := range a {
		fmt.Println(v)
	}
}

type Number interface {
	int
	float32
}

func main() {
	s := []string{"hello", "hi"}
	n := []int{1, 2, 3, 4}
	Print(s)
	Print(n)
}
