package doublylinkedlist

import (
	"fmt"

	"github.com/golanglibs/gocollections/comparer"
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/list"
)

/*
A doubly linked list. Implements Lister and Collectioner.
DoublyLinkedList is not thread safe
*/
type DoublyLinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	equals func(*T, *T) bool
	size   int
}

/*
Creates a new instance of DoublyLinkedList with the given elements with a default equality comparer
and returns pointer to the instance.
If no elements are given, then an empty list is created. Elements must be comparable
*/
func New[K comparable](elements ...K) DoublyLinkedList[K] {
	size := len(elements)

	head, tail := initializeHeadAndTailFromSlice(elements)

	return DoublyLinkedList[K]{
		head:   head,
		tail:   tail,
		equals: comparer.DefaultEquals[K],
		size:   size,
	}
}

/*
Creates a new instance of DoublyLinkedList with the given elements with nil equality comparer
and returns pointer to the instance.
If no elements are given, then an empty list is created. Elements can be of any type
*/
func NewOfAny[T any](elements ...T) DoublyLinkedList[T] {
	head, tail := initializeHeadAndTailFromSlice(elements)

	return DoublyLinkedList[T]{
		head: head,
		tail: tail,
		size: len(elements),
	}
}

func initializeHeadAndTailFromSlice[T any](elements []T) (head *node[T], tail *node[T]) {
	head = newEmptyNode[T]()
	tail = newEmptyNode[T]()
	head.Next = tail
	tail.Prev = head

	node := head
	for _, v := range elements {
		next := node.Next
		element := newNode(v, node, next)

		node.Next = element
		element.Prev = node
		element.Next = next
		next.Prev = element

		node = element
	}

	return head, tail
}

/*
Creates a new instance of OoublyLinkedList from the given collection with a default equality comparer
and returns pointer to the new instance.
Elements of the given collection must be comparable
*/
func NewFromCollection[K comparable](c generic.Collectioner[K]) DoublyLinkedList[K] {
	head, tail := initializeHeadAndTailFromCollection(c)

	return DoublyLinkedList[K]{
		head:   head,
		tail:   tail,
		equals: comparer.DefaultEquals[K],
		size:   c.Size(),
	}
}

/*
Creates a new instance of DoublyLinkedList from the given collection with nil equality comparer
and returns pointer to the instance.
Elements of the given collection can be of any type
*/
func NewOfAnyFromCollection[T any](c generic.Collectioner[T]) DoublyLinkedList[T] {
	head, tail := initializeHeadAndTailFromCollection(c)

	return DoublyLinkedList[T]{
		head: head,
		tail: tail,
		size: c.Size(),
	}
}

func initializeHeadAndTailFromCollection[T any](c generic.Collectioner[T]) (head *node[T], tail *node[T]) {
	head = newEmptyNode[T]()
	tail = newEmptyNode[T]()
	head.Next = tail
	tail.Prev = head

	node := head
	c.ForEach(func(v *T) {
		copiedVal := *v
		next := node.Next
		current := newNode(copiedVal, node, next)
		node.Next = current
		next.Prev = current
		node = current
	})

	return head, tail
}

/*
Sets the equality comparer with the given equals function. Implements Lister.SetEqualityComparer
*/
func (dll *DoublyLinkedList[T]) SetEqualityComparer(equals func(*T, *T) bool) {
	dll.equals = equals
}

/*
Retrieves and returns a reference to the element at the given index. Panics if the given index is out of
range. Iterates through the DoublyLinkedList to find the given index from
either the head or tail depending on whether the index is in the first or last half of the list.
This operation has time complexity of O(n) where n is the number of elements in the list.
Implements lister.At
*/
func (dll *DoublyLinkedList[T]) At(index int) *T {
	if !dll.isValidIndex(index) {
		err := fmt.Sprintf(
			"DoublyLinkedList.At could not retrieve element because given index %d is out of range",
			index,
		)
		panic(err)
	}

	return &dll.findNodeAtIndex(index).Value
}

/*
Returns the reference to the value of the head of the DoublyLinkedList. Panics if the list is empty.
*/
func (dll *DoublyLinkedList[T]) Front() *T {
	if dll.size == 0 {
		panic("DoublyLinkedList.Front failed because the list is empty")
	}

	return &dll.head.Next.Value
}

/*
Returns value of the tail of the DoublyLinkedList. Panics if the list is empty.
*/
func (dll *DoublyLinkedList[T]) Back() *T {
	if dll.size == 0 {
		panic("DoublyLinkedList.Back failed because the list is empty")
	}

	return &dll.tail.Prev.Value
}

/*
Sets the given value at the given index. Panics if the given index is out of range.
Implements lister.Set
*/
func (dll *DoublyLinkedList[T]) Set(index int, value T) {
	if !dll.isValidIndex(index) {
		err := fmt.Sprintf(
			"DoublyLinkedList.Set could not set given value because given index %d is out of range",
			index,
		)
		panic(err)
	}

	node := dll.findNodeAtIndex(index)
	node.Value = value
}

/*
Returns the cached size of the DoublyLinkedList. This operation has time complexity of O(1).
Implements lister.Size and Collectioner.Size
*/
func (dll *DoublyLinkedList[T]) Size() int {
	return dll.size
}

/*
Returns true if the DoublyLinkedList is empty. Implements Lister.Empty and Collectioner.Empty
*/
func (dll *DoublyLinkedList[T]) Empty() bool {
	return dll.size == 0
}

/*
Appends the given element to the tail of the DoublyLinkedList. Always returns true in case of DoublyLinkedList.
Implements Lister.Add and Collectioner.Add
*/
func (dll *DoublyLinkedList[T]) Add(element T) bool {
	newTail := newNode(element, nil, nil)
	prev := dll.tail.Prev

	prev.Next = newTail
	newTail.Prev = prev
	newTail.Next = dll.tail
	dll.tail.Prev = newTail
	dll.size++

	return true
}

/*
Removes the tail of the DoublyLinkedList
*/
func (dll *DoublyLinkedList[T]) RemoveBack() {
	if dll.size == 0 {
		panic("DoublyLinkedList.RemoveBack cannot remove tail because list is empty")
	}

	dll.removeNode(dll.tail.Prev)
}

/*
Inserts the given element at the given index and return true.
If the given index is out of range, DoublyLinkedList.Insert will not add the element and return false.
Implements Lister.Insert
*/
func (dll *DoublyLinkedList[T]) Insert(index int, element T) bool {
	if index < 0 || index > dll.size { // index == dll.size is ok to allow insertion at end after last index
		return false
	}

	newElement := newNode(element, nil, nil)

	nodeAtInsertIndex := dll.findNodeAtIndex(index)
	prev := nodeAtInsertIndex.Prev

	prev.Next = newElement
	newElement.Prev = prev
	newElement.Next = nodeAtInsertIndex
	nodeAtInsertIndex.Prev = newElement

	dll.size++

	return true
}

/*
Prepends the given element to the head of the DoublyLinkedList
*/
func (dll *DoublyLinkedList[T]) AddToFront(element T) {
	newHead := newNode(element, nil, nil)
	next := dll.head.Next

	dll.head.Next = newHead
	newHead.Prev = dll.head
	newHead.Next = next
	next.Prev = newHead
	dll.size++
}

/*
Removes the head of the DoublyLinkedList
*/
func (dll *DoublyLinkedList[T]) RemoveFront() {
	if dll.size == 0 {
		panic("DoublyLinkedList.RemoveFront cannot remove head because list is empty")
	}

	dll.removeNode(dll.head.Next)
}

/*
Removes the first occurrence of the element if found from the DoublyLinkedList and returns true.
If the element is not found, returns false.
Equality is determined by the equality comparer set eith automatically (through constructors for comparable
elements) of manually (through Lister.SetEqualityComparer)
Implements Lister.Remove and Collectioner.Remove
*/
func (dll *DoublyLinkedList[T]) Remove(element T) bool {
	nodeIndex, nodeToRemove := dll.findNode(element)
	if nodeIndex == -1 {
		return false
	}

	dll.removeNode(nodeToRemove)

	return true
}

/*
Removes the element at the given index and returns true. Panics if the given index is out of range
Iterates through the DoublyLinkedList to find the given index from
either the head or tail depending on whether the index is in the first or last half of the list.
Implements Lister.RemoveAt
*/
func (dll *DoublyLinkedList[T]) RemoveAt(index int) {
	if !dll.isValidIndex(index) {
		err := fmt.Sprintf(
			"DoublyLinkedList.RemoveAt cannot remove element at index %d because it is out of range",
			index,
		)
		panic(err)
	}

	nodeToRemove := dll.findNodeAtIndex(index)
	dll.removeNode(nodeToRemove)
}

/*
Iterates through the DoublyLinkedList to find the node at the given index
in the fastest way possible depending on the given index.
If the index to find is in the first half of the list, this method will iterate from the head of the list.
If the index to find is in the last half of the list, this method will iterate from the tail of the list.
This method assumes that the given index is in range of the current DoublyLinkedList
*/
func (dll *DoublyLinkedList[T]) findNodeAtIndex(index int) *node[T] {
	startFromBeginning := dll.size-index > index

	if startFromBeginning {
		current := dll.head.Next
		for i := 0; i < dll.size; i++ {
			if i == index {
				return current
			}

			current = current.Next
		}
	}

	current := dll.tail
	for i := dll.size; i >= 0; i-- {
		if i == index {
			return current
		}

		current = current.Prev
	}

	panic(fmt.Sprintf("Node at index %d should have been found", index))
}

func (dll *DoublyLinkedList[T]) removeNode(node *node[T]) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
	node = nil

	dll.size--
}

/*
Returns the index of the first occurrence of the given value.
Returns -1 if the element is not found.
Equality is determined by the equality comparer set either automatically (through constructors for comparable
elements) of manually (through Lister.SetEqualityComparer)
Implements Lister.IndexOf
*/
func (dll *DoublyLinkedList[T]) IndexOf(element T) int {
	nodeIndex, _ := dll.findNode(element)

	return nodeIndex
}

/*
Returns true if the given element exists in the DoublyLinkedList. Returns false otherwise.
Equality is determined by the equality comparer set either automatically (through constructors for comparable
elements) or manually (through Lister.SetEqualityComparer).
Implements Lister.Contains and Collectioner.Contains
*/
func (dll *DoublyLinkedList[T]) Contains(element T) bool {
	nodeIndex, _ := dll.findNode(element)

	return nodeIndex != -1
}

func (dll *DoublyLinkedList[T]) findNode(element T) (nodeIndex int, node *node[T]) {
	if dll.equals == nil {
		panic("Cannot compute equality of elements since equality comparer is not set")
	}

	current := dll.head.Next
	for i := 0; i < dll.size; i++ {
		if dll.equals(&current.Value, &element) {
			return i, current
		}

		current = current.Next
	}

	return -1, nil
}

/*
Empties the current DoublyLinkedList. Implements Lister.Clear and Collectioner.Clear
*/
func (dll *DoublyLinkedList[T]) Clear() {
	dll.head.Next = dll.tail
	dll.tail.Prev = dll.head
	dll.size = 0
}

/*
Returns a sub list of the current DoublyLinkedList from start index (inclusive) to end index (exclusive).
The return sub list is a new, copied list of the current DoublyLinkedList.
Implements Lister.SubList
*/
func (dll *DoublyLinkedList[T]) SubList(start int, end int) list.Lister[T] {
	if !dll.isValidIndex(start) || end > dll.size || start >= end {
		panic("DoublyLinkedList.SubList Cannot create a sub list because invalid range was given")
	}

	subListHead := newEmptyNode[T]()
	subListTail := newEmptyNode[T]()
	subListHead.Next = subListTail
	subListTail.Prev = subListHead

	node := dll.findNodeAtIndex(start)
	nodeCopy := subListHead
	for i := start; i < end; i++ {
		currentCopy := copyNode(node)
		next := nodeCopy.Next

		nodeCopy.Next = currentCopy
		currentCopy.Prev = nodeCopy
		currentCopy.Next = next
		next.Prev = currentCopy

		node = node.Next
		nodeCopy = currentCopy
	}

	return &DoublyLinkedList[T]{
		head:   subListHead,
		tail:   subListTail,
		equals: dll.equals,
		size:   end - start,
	}
}

func (dll *DoublyLinkedList[T]) isValidIndex(index int) bool {
	return 0 <= index && index < dll.size
}

/*
Loops through the DoublyLinkedList and executes the given "do" function on each element.
Implements Lister.ForEach and Collectioner.ForEach
*/
func (dll *DoublyLinkedList[T]) ForEach(do func(*T)) {
	current := dll.head.Next
	for i := 0; i < dll.size; i++ {
		do(&current.Value)
		current = current.Next
	}
}
