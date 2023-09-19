package queue

type Queuer[T any] interface {
	/* Sets the equality comparer to the given function */
	SetEqualityComparer(equals func(*T, *T) bool)

	/* Returns the size of the queue */
	Size() int

	/* Returns true if the queue is empty. Otherwise, false */
	Empty() bool

	/* Pushes the given value to the back of the queue */
	Enqueue(element T)

	/* Removes the element at the front of the queue. Panics if the queue is empty */
	Dequeue()

	/*
		Returns a reference to the element at the front of the queue without removing it. Panics if the queue
		is empty
	*/
	Peek() *T

	/* Returns true if the given value is found in the queue. Otherwise, false */
	Contains(element T) bool

	/* Empties the queue. Operations performed depends on the implementation */
	Clear()

	/* Iterates through each element in the queue and executes the given function */
	ForEach(do func(*T))
}
