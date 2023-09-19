package generic

type Collectioner[T any] interface {
	/* Returns the size of the collection */
	Size() int

	/* Returns true if the collection is empty. Otherwise, return false */
	Empty() bool

	/*
		Adds the given value to the collection. Returns true if the value was added successfully. Otherwise,
		false
	*/
	Add(element T) bool

	/*
	   Removes the given value from the collection if found. Returns true if the value was found and removed.
	   Otherwise, false
	*/
	Remove(element T) bool

	/* Returns true if the given value was found in the collection. Otherwise, false */
	Contains(element T) bool

	/* Iterates through each element in the collection and executes the given function */
	ForEach(do func(*T))
}
