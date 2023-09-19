package doublylinkedlist

type node[T any] struct {
	Value T
	Prev  *node[T]
	Next  *node[T]
}

func newEmptyNode[T any]() *node[T] {
	return &node[T]{}
}

func newNode[T any](value T, prev *node[T], next *node[T]) *node[T] {
	return &node[T]{
		Value: value,
		Prev:  prev,
		Next:  next,
	}
}

func copyNode[T any](nodeToCopy *node[T]) *node[T] {
	return &node[T]{
		Value: nodeToCopy.Value,
	}
}
