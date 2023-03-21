package stack

// Stack implements a simple, pointer-based stack.
type Stack[T any] struct {
	elements *stackElement[T]
}

// stackElements holds a single data item, and pointer to the next element.
type stackElement[T any] struct {
	data T
	next *stackElement[T]
}

// New creates & returns a new *Stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds a data item to the Stack.
func (s *Stack[T]) Push(data T) {
	newElement := stackElement[T]{
		data: data,
		next: s.elements,
	}
	s.elements = &newElement
}

// Pop removes (and returns) the most recently added item from the Stack.
func (s *Stack[T]) Pop() T {
	if s.Empty() {
		var data T // make a 0-value of type T
		return data
	}
	top := s.elements.data
	s.elements = s.elements.next
	return top
}

// Empty returns true IFF there are no data items in the Stack.
func (s *Stack[T]) Empty() bool {
	return s.elements == nil
}
