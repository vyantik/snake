package deque

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type Deque[T any] struct {
	front *Node[T]
	back  *Node[T]
	size  int
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (d *Deque[T]) PushFront(value T) {
	newNode := &Node[T]{value: value}
	if d.size == 0 {
		d.front = newNode
		d.back = newNode
	} else {
		newNode.next = d.front
		d.front.prev = newNode
		d.front = newNode
	}
	d.size++
}

func (d *Deque[T]) PushBack(value T) {
	newNode := &Node[T]{value: value}
	if d.size == 0 {
		d.front = newNode
		d.back = newNode
	} else {
		newNode.prev = d.back
		d.back.next = newNode
		d.back = newNode
	}
	d.size++
}

func (d *Deque[T]) PopFront() T {
	var zero T
	if d.size == 0 {
		return zero
	}
	value := d.front.value
	if d.size == 1 {
		d.front = nil
		d.back = nil
	} else {
		d.front = d.front.next
		d.front.prev = nil
	}
	d.size--
	return value
}

func (d *Deque[T]) PopBack() T {
	var zero T
	if d.size == 0 {
		return zero
	}
	value := d.back.value
	if d.size == 1 {
		d.front = nil
		d.back = nil
	} else {
		d.back = d.back.prev
		d.back.next = nil
	}
	d.size--
	return value
}

func (d *Deque[T]) Front() T {
	var zero T
	if d.size == 0 {
		return zero
	}
	return d.front.value
}

func (d *Deque[T]) Back() T {
	var zero T
	if d.size == 0 {
		return zero
	}
	return d.back.value
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		current := d.front
		for current != nil {
			ch <- current.value
			current = current.next
		}
		close(ch)
	}()
	return ch
}

func ElementInDeque[T comparable](d *Deque[T], element T) bool {
	for v := range d.Iter() {
		if v == element {
			return true
		}
	}
	return false
}

func Vector2InDeque(d *Deque[rl.Vector2], element rl.Vector2) bool {
	for v := range d.Iter() {
		if rl.Vector2Equals(v, element) {
			return true
		}
	}
	return false
}
