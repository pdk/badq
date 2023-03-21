package queue

import "github.com/pdk/badq/stack"

// Queue implements a standard queue using 2 stacks.
type Queue[T any] struct {
	readStack *stack.Stack[T]
	holdStack *stack.Stack[T]
}

// New creates & returns a new *Queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{
		readStack: stack.New[T](),
		holdStack: stack.New[T](),
	}
}

// Push adds a data item to the Queue.
func (q *Queue[T]) Push(data T) {

	for !q.readStack.Empty() {
		q.holdStack.Push(q.readStack.Pop())
	}

	q.readStack.Push(data)
}

// Pull removes (and returns) the oldest value from the Queue.
func (q *Queue[T]) Pull() T {

	for !q.holdStack.Empty() {
		q.readStack.Push(q.holdStack.Pop())
	}

	return q.readStack.Pop()
}

// Empty returns true IFF there are no items in the Queue.
func (q *Queue[T]) Empty() bool {
	return q.readStack.Empty() && q.holdStack.Empty()
}
