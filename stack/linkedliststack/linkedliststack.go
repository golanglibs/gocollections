package linkedliststack

import (
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/list/doublylinkedlist"
)

/*
Linked list based stack. Last element to be pushed will be popped first (LIFO).
It uses gocollections/list/doublylinkedlist to perform stack operations. This implementation will generally
perform slower than Array-based stack (gocollections/list/arraystack) due to being pointer based but will
be more efficient in terms of memory usage if the size of the stack fluctuates greatly since any removed
elements will be garbage-collected
Implements Stacker and Collectioner.
LinkedListStack is not thread safe
*/
type LinkedListStack[T any] struct {
	container doublylinkedlist.DoublyLinkedList[T]
}

/*
Creates a new instance of Stack with the given elements with a default equality comparer
and returna pointer to the instance.
If no elements are given, then an empty stack is created. Elements must be comparable
*/
func New[K comparable](elements ...K) LinkedListStack[K] {
	return LinkedListStack[K]{
		container: doublylinkedlist.New(elements...),
	}
}

/*
Creates a new instance of Stack with the given elements with nil equality comparer
and returns pointer to the instance.
If no elements are given, then an empty queue is created. Elements can be of any type
*/
func NewOfAny[T any](elements ...T) LinkedListStack[T] {
	return LinkedListStack[T]{
		container: doublylinkedlist.NewOfAny(elements...),
	}
}

/*
Creates a new instance of Stack from the given collection with a default equality comparer
and returns pointer to the new instance.
Elements of the given collection must be comparable
*/
func NewFromCollection[K comparable](c generic.Collectioner[K]) LinkedListStack[K] {
	return LinkedListStack[K]{
		container: doublylinkedlist.NewFromCollection(c),
	}
}

/*
Creates a new instance of Stack from the given collection with nil equality comparer
and returns pointer to the instance.
Elements of the given collection can be of any type
*/
func NewOfAnyFromCollection[T any](c generic.Collectioner[T]) LinkedListStack[T] {
	return LinkedListStack[T]{
		container: doublylinkedlist.NewOfAnyFromCollection(c),
	}
}

/*
Sets the equality comparer with the given equals function. Implements Stacker.SetEqualityComparer
*/
func (s *LinkedListStack[T]) SetEqualityComparer(equals func(a *T, b *T) bool) {
	s.container.SetEqualityComparer(equals)
}

/*
Returns the length of the Stack. Implements Stacker.Size and Collectioner.Size
*/
func (s *LinkedListStack[T]) Size() int {
	return s.container.Size()
}

/*
Returns true if the Stack is empty. Implements Stacker.Empty and Collectioner.Empty
*/
func (s *LinkedListStack[T]) Empty() bool {
	return s.container.Empty()
}

/*
Adds the given element to the stack. Implements Stacker.Push
*/
func (s *LinkedListStack[T]) Push(element T) {
	s.container.Add(element)
}

/*
Removes the most recently pushed element in the stack. Panics if Stack is empty.
Implements Stacker.Pop
*/
func (s *LinkedListStack[T]) Pop() {
	if s.container.Empty() {
		panic("Stack.Pop failed because stack is empty")
	}

	s.container.RemoveBack()
}

/*
Returns a reference to the most recently pushed element in the stack without removing it. Panics if Stack is
empty. Implements Stacker.Peek
*/
func (s *LinkedListStack[T]) Peek() *T {
	if s.container.Empty() {
		panic("Stack.Peek failed because stack is empty")
	}

	return s.container.Back()
}

/*
Adds the given element to the front of the stack. Always returns true.
Stack.Add functions exactly the same as Stack.Push except that it returns bool.
Implements Stacker.Add and Collectioner.Add
*/
func (s *LinkedListStack[T]) Add(element T) bool {
	s.container.Add(element)
	return true
}

/*
Removes the the given element and returns true if present in the Stack.
Returns false if the given element does not exist.
Implements Stacker.Remove and Collectioner.Remove
*/
func (s *LinkedListStack[T]) Remove(element T) bool {
	return s.container.Remove(element)
}

/*
Returns true if the given element exists in the Stack. Returns false otherwise.
Implements Stacker.Contains and Collectioner.Contains
*/
func (s *LinkedListStack[T]) Contains(element T) bool {
	return s.container.Contains(element)
}

/*
Empties the Stack.
Implements Stacker.Clear and Collectioner.Clear
*/
func (s *LinkedListStack[T]) Clear() {
	s.container.Clear()
}

/*
Iterates through each element in the stack and executes the given function. Note that the order of
iteration will be the opposite of the order each element would be popped
Implements Stacker.ForEach and Collectioner.ForEach
*/
func (s *LinkedListStack[T]) ForEach(do func(*T)) {
	s.container.ForEach(do)
}
