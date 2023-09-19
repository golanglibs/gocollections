package arraylist

import (
	"fmt"

	"github.com/golanglibs/gocollections/comparer"
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/list"
)

/*
Array-based list. Uses go's native slice internally.
Implements Lister and Collectioner.
List is not thread safe
*/
type List[T any] struct {
	equals    func(*T, *T) bool
	container []T
	size      int
	cap       int
}

/*
Creates a new instance of List with the given elements with a default equality comparer
and returns it
If no elements are given, then an empty list is created. Elements must be comparable
*/
func New[K comparable](elements ...K) List[K] {
	var container []K

	size := len(elements)
	if size > 0 {
		container = make([]K, size)
		copy(container, elements)
	}

	return List[K]{
		equals:    comparer.DefaultEquals[K],
		container: container,
		size:      size,
		cap:       size,
	}
}

/*
Creates a new instance of List with the given elements with nil equality comparer and returns it
If no elements are given, then an empty list is created. Elements can be of any type
*/
func NewOfAny[T any](elements ...T) List[T] {
	var container []T

	size := len(elements)
	if size > 0 {
		container = append(container, elements...)
	}

	return List[T]{
		container: container,
		size:      size,
		cap:       size,
	}
}

/*
Creates a new instance of List from the given collection with a default equality comparer and returns it
Elements of the given collection must be comparable
*/
func NewFromCollection[K comparable](c generic.Collectioner[K]) List[K] {
	var copiedSlice []K
	c.ForEach(func(element *K) {
		copiedSlice = append(copiedSlice, *element)
	})
	size := len(copiedSlice)

	return List[K]{
		equals:    comparer.DefaultEquals[K],
		container: copiedSlice,
		size:      size,
		cap:       size,
	}
}

/*
Creates a new instance of List from the given collection with nil equality comparer and returns it
Elements of the given collection can be of any type
*/
func NewOfAnyFromCollection[T any](c generic.Collectioner[T]) List[T] {
	var copiedSlice []T
	c.ForEach(func(element *T) {
		copiedSlice = append(copiedSlice, *element)
	})
	size := len(copiedSlice)

	return List[T]{
		container: copiedSlice,
		size:      size,
		cap:       size,
	}
}

/*
Sets the equality comparer with the given equals function. Implements Lister.SetEqualityComparer
*/
func (l *List[T]) SetEqualityComparer(equals func(*T, *T) bool) {
	l.equals = equals
}

/*
Retrieves and returns a reference to the element at the given index. Panics if the given index is out of range.
Implements lister.At
*/
func (l *List[T]) At(index int) *T {
	if !l.isValidIndex(index) {
		err := fmt.Sprintf("List.At could not retrieve element because given index %d is out of range", index)
		panic(err)
	}

	return &l.container[index]
}

/*
Sets the given value at the given index. Panics if the given index is out of range.
Implements lister.Set
*/
func (l *List[T]) Set(index int, value T) {
	if !l.isValidIndex(index) {
		err := fmt.Sprintf("List.Set could not set given value because given index %d is out of range", index)
		panic(err)
	}

	l.container[index] = value
}

/*
Returns the length of the List. Implements lister.Size
*/
func (l *List[T]) Size() int {
	return l.size
}

/*
Returns true if the List is empty. Implements Lister.Empty and Collectioner.Empty
*/
func (l *List[T]) Empty() bool {
	return l.size == 0
}

/*
Returns a reference to the first element (at index 0) in the list
*/
func (l *List[T]) Front() *T {
	if l.size == 0 {
		panic("ArrayList.Front failed because the list is empty")
	}

	return &l.container[0]
}

func (l *List[T]) Back() *T {
	if l.size == 0 {
		panic("ArrayList.Back failed because the list is empty")
	}

	return &l.container[l.size-1]
}

/*
Adds the given element to the end of the List. Add will always return true for List.
Implements Lister.Add and Collectioner.Add
*/
func (l *List[T]) Add(element T) bool {
	if l.size == l.cap {
		l.container = append(l.container, element)
		l.cap++
	} else {
		l.container[l.size] = element
	}

	l.size++
	return true
}

/*
Removes the last element of the list.
Implements Lister.RemoveBack
*/
func (l *List[T]) RemoveBack() {
	if l.size == 0 {
		panic("ArrayList.RemoveBack failed because the list is empty")
	}

	l.size--
}

/*
Adds the given element at the given index and return true.
If the given index is out of range, List.Insert will not add the element and return false.
Implements Lister.Insert
*/
func (l *List[T]) Insert(index int, element T) bool {
	if index < 0 || index > l.size { // index == length is ok to allow for insertion at end after last index
		return false
	}

	if index == l.size {
		l.Add(element)
		return true
	}

	l.container = append(l.container[:index+1], l.container[index:]...)
	l.container[index] = element
	l.size++
	if l.size > l.cap {
		l.cap++
	}

	return true
}

/*
Adds the given value to the front of the list.
Implements Lister.AddToFront
*/
func (l *List[T]) AddToFront(element T) {
	l.Insert(0, element)
}

/*
Removes the first element of the list
Implements Lister.RemoveFront
*/
func (l *List[T]) RemoveFront() {
	l.RemoveAt(0)
}

/*
Removes the first occurrence of the given element from the List.
Returns true if the element is found and removed. Returns false if the element to remove is not found.
Equality is determined by the equality comparer set either automatically (through constructors for comparable
elements or manually (through Lister.SetEqualityComparer).
Panics if the equality comparer "equals" method is not set.
Implements Lister.Remove and Collectioner.Remove
*/
func (l *List[T]) Remove(element T) bool {
	elementIndex := l.indexOf(element)
	if elementIndex == -1 {
		return false
	}

	l.container = append(l.container[:elementIndex], l.container[elementIndex+1:]...)
	l.size--

	return true
}

/*
Removes the element at the given index from the List. Panics if the given index is out of range
Implements Lister.RemoveAt
*/
func (l *List[T]) RemoveAt(index int) {
	if !l.isValidIndex(index) {
		err := fmt.Sprintf(
			"ArrayList.RemoveAt cannot remove element at index %d because it is out of range",
			index,
		)
		panic(err)
	}

	l.container = append(l.container[:index], l.container[index+1:]...)
	l.size--
}

/*
Returns the index of the given element.
Returns -1 if the element is not found.
Equality is determined by the equality comparer set either automatically (through constructors for comparable elements)
or manually (through Lister.SetEqualityComparer).
Implements Lister.IndexOf
*/
func (l *List[T]) IndexOf(element T) int {
	return l.indexOf(element)
}

/*
Returns true if the given element exists in the List. Returns false otherwise.
Equality is determined by the equality comparer set either automatically (through constructors for comparable elements)
or manually (through Lister.SetEqualityComparer).
Implements Lister.Contains and Collectioner.Contains
*/
func (l *List[T]) Contains(element T) bool {
	return l.indexOf(element) != -1
}

func (l *List[T]) indexOf(element T) int {
	if l.equals == nil {
		panic("Cannot compute equality of elements since equality comparer is not set")
	}

	for i, v := range l.container {
		if l.equals(&v, &element) {
			return i
		}
	}

	return -1
}

/*
Returns a sub list of the current List from start index (inclusive) to end index (exclusive).
The returned sub list is a new, copied list of the current List.
Implements Lister.SubList
*/
func (l *List[T]) SubList(start int, end int) list.Lister[T] {
	if !l.isValidIndex(start) || end > len(l.container) || start >= end {
		panic("List.SubList Cannot create a sub list because invalid range was given")
	}

	newSize := end - start
	newContainer := make([]T, newSize)
	copy(newContainer, l.container[start:end])

	return &List[T]{
		container: newContainer,
		size:      newSize,
		cap:       newSize,
	}
}

func (l *List[T]) isValidIndex(index int) bool {
	return 0 <= index && index < l.size
}

/*
Empties the current List. This does not actually set the internal slice to nil. It simply sets the internal
size counter to zero.
Implements Lister.Clear and Collectioner.Clear
*/
func (l *List[T]) Clear() {
	l.size = 0
}

/*
Loops through the List and executes the given "do" function on each element.
Implements Seter.ForEach and Collectioner.ForEach
*/
func (l *List[T]) ForEach(do func(*T)) {
	for _, v := range l.container {
		do(&v)
	}
}
