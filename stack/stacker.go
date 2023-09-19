package stack

type Stacker[T any] interface {
	/* Sets the equality comparer to the given function */
	SetEqualityComparer(equals func(*T, *T) bool)

	/* Returns the size of the stack */
	Size() int

	/* Returns true if the stack is empty. Otherwise, false */
	Empty() bool

	/* Pushes the given value to the stack */
	Push(element T)

	/* Removes the element to the stack. Panics if the stack is empty */
	Pop()

	/*
		Returns a reference to the recently pushed element in the stack without removing it. Panics if stack is
		empty
	*/
	Peek() *T

	/* Returns true if the given value is found in the stack. Otherwise, false */
	Contains(element T) bool

	/* Empties the stack. Operations performed depends on the implementation */
	Clear()

	/*
		Iterates through each element in the stack and executes the given function. Note that the order of
		iteration will be the opposite of the order each element would be popped
	*/
	ForEach(do func(*T))
}
