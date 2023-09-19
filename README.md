# gocollections
Basic `Generic` Collections Library for Golang with interfaces to provide implementation-agnostic abstractions 
over various collections

## List of Implemented Data Structures
* [ArrayList](./list/arraylist/list.go)
* [DoublyLinkedList](./list/doublylinkedlist/doublylinkedlist.go)
* [HashSet](./set/hashset/set.go)
* [LinkedListQueue](./queue/linkedlistqueue/queue.go)
* [PriorityQueue](./queue/priorityqueue/pq.go)
* [ArrayStack](./stack/arraystack/stack.go)
* [LinkedListStack](./stack/linkedliststack/linkedliststack.go)

## Provided Collection Interfaces and their implementations
* [Collectioner[T any]](./generic/collectioner.go)
    * Encompasses all collections implemented in this library
    * Provides the following operations:
        * `Size() int`
        * `Empty() bool`
        * `Add(element T) bool`
        * `Remove(element T) bool`
        * `Contains(element T) bool`
        * `ForEach(do func(*T))`
    * Implemented By:
        * [ArrayList](./list/arraylist/list.go)
        * [DoublyLinkedList](./list/doublylinkedlist/doublylinkedlist.go)
        * [HashSet](./set/hashset/set.go)
        * [LinkedListQueue](./queue/linkedlistqueue/queue.go)
        * [PriorityQueue](./queue/priorityqueue/pq.go)
        * [ArrayStack](./stack/arraystack/stack.go)
        * [LinkedListStack](./stack/linkedliststack/linkedliststack.go)

* [Lister[T any]](./list/lister.go)
    * Provides operations for list-like collections
    * Provides the following operations:
        * `SetEqualityComparer(equals func(*T, *T) bool)`:
        * `At(index int) *T`
        * `Set(index int, value T)`
        * `Size() int`
        * `Empty() bool`
        * `Front() *T`
        * `Back() *T`
        * `Add(element T) bool`
        * `RemoveBack()`
        * `Insert(index int, value T) bool`
        * `AddToFront(element T)`
        * `RemoveFront()`
        * `Remove(element T) bool`
        * `RemoveAt(index int)`
        * `IndexOf(element T) int`
        * `Contains(element T) bool`
        * `SubList(start int, end int) Lister[T]`
        * `Clear()`
        * `ForEach(do func(*T))`
    * Implemented by:   
        * [ArrayList](./list/arraylist/list.go)
        * [DoublyLinkedList](./list/doublylinkedlist/doublylinkedlist.go)

* [Seter[K comparable]](./set/seter.go)
    * Provides operations for set-like collections
    * Provides the following operations:
        * `Size() int`
        * `Empty() bool`
        * `Add(element K) bool`
        * `Remove(element K) bool`
        * `Contains(element K) bool`
        * `Equals(set Seter[K]) bool`
        * `Intersects(set Seter[K]) bool`
        * `GetIntersection(set Seter[K]) Seter[K]`
        * `GetUnion(set Seter[K]) Seter[K]`
        * `IsSupersetOf(set Seter[K]) bool`
        * `IsSubsetOf(set Seter[K]) bool`
        * `Clear()`
        * `ForEach(do func(*K))`
    * Implemented By:
        * [HashSet](./set/hashset/set.go)

* [Queuer[T any]](./queue/queuer.go)
    * Provides operations for queue-like collections
    * Provides the following operations
	    * `SetEqualityComparer(equals func(*T, *T) bool)`
	    * `Size() int`
	    * `Empty() bool`
	    * `Enqueue(element T)`
	    * `Dequeue()`
	    * `Peek() *T`
	    * `Contains(element T) bool`
	    * `Clear()`
	    * `ForEach(do func(*T))`
    * Implemented By:
        * [LinkedListQueue](./queue/linkedlistqueue/queue.go)
        * [PriorityQueue](./queue/priorityqueue/pq.go) - Binary Heap

* [Stacker[T any]](./stack/stacker.go)
    * Provides operations for stack-like collections
    * Provides the following operations:
	    * `SetEqualityComparer(equals func(*T, *T) bool)`
	    * `Size() int`
	    * `Empty() bool`
	    * `Push(T)`
	    * `Pop()`
	    * `Peek() *T`
	    * `Contains(element T) bool`
	    * `Clear()`
	    * `ForEach(do func(*T))`
    * Implemented By:
        * [ArrayStack](./stack/arraystack/stack.go)
        * [LinkedListStack](./stack/linkedliststack/linkedliststack.go)

## Possible Improvements
* Add `SortedMap (TreeMap)` and `SortedSet (TreeSet)`
* Add `Stream APIs` using `Collectioner`
* Add more collection methods such as `Map` and `Filter`
