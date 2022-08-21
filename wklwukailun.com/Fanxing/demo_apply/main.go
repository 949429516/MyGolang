package main

import "fmt"

type NewType[T any] struct {
}

func (n NewType[T]) get(in T) (out T) {
	return in
}
func main() {
	n := NewType[int]{}
	f := n.get(1)
	fmt.Println(f)
}
