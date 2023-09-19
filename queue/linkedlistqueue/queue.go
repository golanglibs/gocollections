package linkedlistqueue

import (
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/list/doublylinkedlist"
)

/*
Linked list based queue. First element to be enqueued will be dequeued first (FIFO).
It uses gocollections/list/doublylinkedlist to perform queue operations.
Implements Queuer and Collectioner.
Queue is not thread safe
*/
type Queue[T any] struct {
	container doublylinkedlist.DoublyLinkedList[T]
}

/*
Creates a new instance of Queue with the given elements with a default equality comparer
and returns pointer to the instance.
If no elements are given, then an empty queue is created. Elements must be comparable
*/
func New[K comparable](elements ...K) Queue[K] {
	return Queue[K]{
		container: doublylinkedlist.New(elements...),
	}
}

/*
Creates a new instance of Queue with the given elements with nil equality comparer
and returns pointer to the instance.
If no elements are given, then an empty queue is created. Elements can be of any type
*/
func NewOfAny[T any](elements ...T) Queue[T] {
	return Queue[T]{
		container: doublylinkedlist.NewOfAny(elements...),
	}
}

/*
Creates a new instance of Queue from the given collection with a default equality comparer
and returns pointer to the new instance.
Elements of the given collection must be comparable
*/
func NewFromCollection[K comparable](c generic.Collectioner[K]) Queue[K] {
	return Queue[K]{
		container: doublylinkedlist.NewFromCollection(c),
	}
}

/*
Creates a new instance of Queue from the given collection with nil equality comparer
and returns pointer to the instance.
Elements of the given collection can be of any type
*/
func NewOfAnyFromCollection[T any](c generic.Collectioner[T]) Queue[T] {
	return Queue[T]{
		container: doublylinkedlist.NewOfAnyFromCollection(c),
	}
}

/*
Sets the equality comparer with the given equals function. Implements Queuer.SetEqualityComparer
*/
func (q *Queue[T]) SetEqualityComparer(equals func(a *T, b *T) bool) {
	q.container.SetEqualityComparer(equals)
}

/*
Returns the length of the Queue. Implements Queuer.Size and Collectioner.Size
*/
func (q *Queue[T]) Size() int {
	return q.container.Size()
}

/*
Returns true if the Queue is empty. Implements Queuer.Empty and Collectioner.Empty
*/
func (q *Queue[T]) Empty() bool {
	return q.container.Empty()
}

/*
Pushes the given value to the back of the queue. Implements Queuer.Enqueue
*/
func (q *Queue[T]) Enqueue(element T) {
	q.container.Add(element)
}

/*
Removes the element at the front of the queue. Panics if Queue is empty.
Implements Queuer.Dequeue
*/
func (q *Queue[T]) Dequeue() {
	if q.container.Empty() {
		panic("Queue.Dequeue failed because queue is empty")
	}

	q.container.RemoveFront()
}

/*
Returns a reference to the element at the front of the queue without removing it. Panics if Queue is empty.
Implements Queuer.Peek
*/
func (q *Queue[T]) Peek() *T {
	if q.container.Empty() {
		panic("Queue.Peek failed because queue is empty")
	}

	return q.container.Front()
}

/*
Adds the given element to the back of the queue. Always returns true.
Queue.Add functions exactly the same as Queue.Enqueue except that it returns bool.
Implements Queuer.Add and Collectioner.Add
*/
func (q *Queue[T]) Add(element T) bool {
	return q.container.Add(element)
}

/*
Removes the the given element and returns true if present in the Queue.
Returns false if the given element does not exist.
Implements Queuer.Remove and Collectioner.Remove
*/
func (q *Queue[T]) Remove(element T) bool {
	return q.container.Remove(element)
}

/*
Returns true if the given element exists in the Queue. Returns false otherwise.
Implements Queuer.Contains and Collectioner.Contains
*/
func (q *Queue[T]) Contains(element T) bool {
	return q.container.Contains(element)
}

/*
Empties the Queue.
Implements Queuer.Clear and Collectioner.Clear
*/
func (q *Queue[T]) Clear() {
	q.container.Clear()
}

/*
Iterates through the Queue and executes the given "do" function on each element.
Implements Queuer.ForEach and Collectioner.ForEach
*/
func (q *Queue[T]) ForEach(do func(*T)) {
	q.container.ForEach(do)
}
