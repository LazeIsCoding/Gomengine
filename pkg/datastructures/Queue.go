package datastructures

/**
Implementation einer Queue mittels Nodes.
Stellt Funktionen folgende Funktionen bereit:

NewQueue[T comparable]()*Queue[T] -> Erzeugt eine neue, leere Queue

(q *queue[T])Enqueue(key T) -> Hängt ein Element an das Ende der Queue an

(q *Queue[T])Dequeue() *Node[T] -> Entfernt ein Element von der Liste
*/

import "fmt"

type Queue[T comparable] struct {
	start *Node[T]
	end   *Node[T]
	size  int
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		start: nil,
		end:   nil,
		size:  0,
	}
}
func (q *Queue[T]) Enqueue(key T) {
	tmp := new(Node[T])
	tmp.key = key
	tmp.next = q.start
	q.start = tmp
	if q.size == 0 {
		q.end = tmp
	}
	q.size++
}

func (q *Queue[T]) Dequeue() *T {
	if q.IsEmpty() {
		return nil
	}
	if q.size == 1 {
		key := q.start.key
		q.start = nil
		q.end = nil
		q.size--
		return &key
	} else {
		tmp := q.start
		for tmp.next != q.end {
			tmp = tmp.next
		}
		tmp.next = nil
		//Delete Node on end?
		key := q.end.key
		q.end = tmp
		q.size--
		return &key
	}

}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Find(key T) *Node[T] {
	tmp := q.start
	for i := 0; tmp.next != nil; i++ {
		if tmp.key == key {
			return tmp
		} else {
			tmp = tmp.next
		}
	}
	return nil
}

func (q *Queue[T]) PrintQueue() {
	if q.IsEmpty() {
		fmt.Println("Empty")
	} else {
		tmp := q.start
		for tmp != nil {
			fmt.Print(tmp.key, "->")
			tmp = tmp.next
		}
		fmt.Println("")
	}
}
