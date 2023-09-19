package arraystack

import (
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/list/arraylist"
)

/*
Array-based stack. Last element to be pushed will be popped first (LIFO).
It uses gocollections/list/arraylist to perform stack operations. This implementation will outperform the
linked list stack implementation (gocollections/list/linkedliststack) in terms of latency due to the
contiguous nature of the internal container which improves CPU cache utilization. However, this comes at the
cost of extra memory usage if there is a big fluctuation in the size of the stack between operations because
it will not free the memory in the internal container that was allocated when there were more elements pushed
in the stack
Implements Stacker and Collectioner.
Stack is not thread safe
*/
type Stack[T any] struct {
	container arraylist.List[T]
}

/*
Creates a new instance of Stack with the given elements with a default equality comparer
and returna pointer to the instance.
If no elements are given, then an empty stack is created. Elements must be comparable
*/
func New[K comparable](elements ...K) Stack[K] {
	return Stack[K]{
		container: arraylist.New(elements...),
	}
}

/*
Creates a new instance of Stack with the given elements with nil equality comparer
and returns pointer to the instance.
If no elements are given, then an empty queue is created. Elements can be of any type
*/
func NewOfAny[T any](elements ...T) Stack[T] {
	return Stack[T]{
		container: arraylist.NewOfAny(elements...),
	}
}

/*
Creates a new instance of Stack from the given collection with a default equality comparer
and returns pointer to the new instance.
Elements of the given collection must be comparable
*/
func NewFromCollection[K comparable](c generic.Collectioner[K]) Stack[K] {
	return Stack[K]{
		container: arraylist.NewFromCollection(c),
	}
}

/*
Creates a new instance of Stack from the given collection with nil equality comparer
and returns pointer to the instance.
Elements of the given collection can be of any type
*/
func NewOfAnyFromCollection[T any](c generic.Collectioner[T]) Stack[T] {
	return Stack[T]{
		container: arraylist.NewOfAnyFromCollection(c),
	}
}

/*
Sets the equality comparer with the given equals function. Implements Stacker.SetEqualityComparer
*/
func (s *Stack[T]) SetEqualityComparer(equals func(a *T, b *T) bool) {
	s.container.SetEqualityComparer(equals)
}

/*
Returns the length of the Stack. Implements Stacker.Size and Collectioner.Size
*/
func (s *Stack[T]) Size() int {
	return s.container.Size()
}

/*
Returns true if the Stack is empty. Implements Stacker.Empty and Collectioner.Empty
*/
func (s *Stack[T]) Empty() bool {
	return s.container.Empty()
}

/*
Adds the given element to the stack. Implements Stacker.Push
*/
func (s *Stack[T]) Push(element T) {
	s.container.Add(element)
}

/*
Removes the most recently pushed element in the stack. Panics if Stack is empty.
Implements Stacker.Pop
*/
func (s *Stack[T]) Pop() {
	if s.container.Empty() {
		panic("Stack.Pop failed because stack is empty")
	}

	s.container.RemoveBack()
}

/*
Returns a reference to the most recently pushed element in the stack without removing it. Panics if Stack is
empty. Implements Stacker.Peek
*/
func (s *Stack[T]) Peek() *T {
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
func (s *Stack[T]) Add(element T) bool {
	s.container.Add(element)
	return true
}

/*
Removes the the given element and returns true if present in the Stack.
Returns false if the given element does not exist.
Implements Stacker.Remove and Collectioner.Remove
*/
func (s *Stack[T]) Remove(element T) bool {
	return s.container.Remove(element)
}

/*
Returns true if the given element exists in the Stack. Returns false otherwise.
Implements Stacker.Contains and Collectioner.Contains
*/
func (s *Stack[T]) Contains(element T) bool {
	return s.container.Contains(element)
}

/*
Empties the Stack.
Implements Stacker.Clear and Collectioner.Clear
*/
func (s *Stack[T]) Clear() {
	s.container.Clear()
}

/*
Iterates through each element in the stack and executes the given function. Note that the order of
iteration will be the opposite of the order each element would be popped
Implements Stacker.ForEach and Collectioner.ForEach
*/
func (s *Stack[T]) ForEach(do func(*T)) {
	s.container.ForEach(do)
}
