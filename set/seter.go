package set

type Seter[K comparable] interface {
	/* Returns the size of the set */
	Size() int

	/* Returns true if the set is empty. Otherwise, false */
	Empty() bool

	/* Adds the given element to the Set. */
	Add(element K) bool

	/* Removes the given element from the Set. */
	Remove(element K) bool

	/* Returns true when the given element exists in the Set. */
	Contains(element K) bool

	/* Returns true when the given set has the equal members as the current set */
	Equals(set Seter[K]) bool

	/* Returns true when the given set has common members with the current set. */
	Intersects(set Seter[K]) bool

	/* Returns a new instance of set with the common members between the current set and the given set. */
	GetIntersection(set Seter[K]) Seter[K]

	/* Returns a new instance of Set with all the members of both the current set and the given set. */
	GetUnion(set Seter[K]) Seter[K]

	/* Returns true if the current set contains all the members of the given set. */
	IsSupersetOf(set Seter[K]) bool

	/* Returns true if the given set has all the members of the current set. */
	IsSubsetOf(set Seter[K]) bool

	/* Empties the set. Operations performed depends on the implementation */
	Clear()

	/* Iterates through each element in the set and executes the given function */
	ForEach(do func(*K))
}
