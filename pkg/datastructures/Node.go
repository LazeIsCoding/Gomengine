package datastructures

type Node[T comparable] struct {
	next *Node[T]
	key  T
}
