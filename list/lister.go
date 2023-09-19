package list

type Lister[T any] interface {
	/* Sets the equality comparer to the given function */
	SetEqualityComparer(equals func(*T, *T) bool)

	/* Retrieves and returns a reference to the element at the given index */
	At(index int) *T

	/* Sets the given value at the given index */
	Set(index int, value T)

	/* Returns the size of the List */
	Size() int

	/* Returns true if the list empty. Otherwise, false */
	Empty() bool

	/* Returns a reference to the value at the front of the list */
	Front() *T

	/* Returns a reference to the value at the back of the list */
	Back() *T

	/* Appends the given value to the back of the list */
	Add(element T) bool

	/* Removes the last element of the list */
	RemoveBack()

	/* Inserts the given value at the given index */
	Insert(index int, value T) bool

	/* Adds the given value to the front of the list */
	AddToFront(element T)

	/* Removes the first element of the list */
	RemoveFront()

	/*
		Removes the first occurrence of the element if found. Returns true if the given element was found and
		removed. Otherwise, false
	*/
	Remove(element T) bool

	/* Removes the element at the given index. Panics if the given index is out of range */
	RemoveAt(index int)

	/*
		Returns the index of the first occurrence of the given value. Returns -1 if the element is not found.
		Equality is determined by the equality comparer set either automatically (because the types of the
		the elements are "comparable") or set manually (through Lister.SetEqualityComparer)
	*/
	IndexOf(element T) int

	/* Returns true if the given value exists in the list. Otherwise, false */
	Contains(element T) bool

	/*
		Returns a sub list of the current list from "start" index (inclusive) to "end" index (exclusive).
		The returned sub list is a new, copied list of the currrent list.
	*/
	SubList(start int, end int) Lister[T]

	/* Empties the list. How the emptying is performed depends on the implementation */
	Clear()

	/* Iterates through each element in the list and executes the given function */
	ForEach(do func(*T))
}
