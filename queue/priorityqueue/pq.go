package priorityqueue

/*
Binary Heap. It uses gocollections/list/arraylist to perform operations
Implements Queuer and Collectioner
"SetEqualityComparer" method is required for "Remove" and "Contains" methods to work properly
PriorityQueue is not thread safe
*/
type PriorityQueue[T any] struct {
	equals    func(*T, *T) bool
	compare   func(*T, *T) bool
	container []T
	cap       int
	size      int
}

func siftUp[T any](heap []T, size int, compare func(*T, *T) bool) {
	i := size
	for i > 1 {
		parent := i >> 1
		if compare(&heap[i], &heap[parent]) {
			heap[i], heap[parent] = heap[parent], heap[i]
		} else {
			break
		}

		i = parent
	}
}

func siftDown[T any](start int, heap []T, size int, compare func(*T, *T) bool) {
	i := start
	for (i << 1) <= size {
		child := i << 1
		if child+1 <= size && compare(&heap[child+1], &heap[child]) {
			child++
		}

		if compare(&heap[child], &heap[i]) {
			heap[child], heap[i] = heap[i], heap[child]
		} else {
			break
		}

		i = child
	}
}

/*
Initializes a new instance of empty PriorityQueue and returns it. "compare" function must be passed in order
to compare elements. If the "compare(e0, e1)" returns true, it means "e0" will have higher priority than "e1"
and will appear earlier than "e1" when dequeueing or popping the top elements
*/
func New[T any](compare func(*T, *T) bool) PriorityQueue[T] {
	var filler T
	return PriorityQueue[T]{
		compare:   compare,
		container: []T{filler},
		cap:       0,
		size:      0,
	}
}

/*
Makes a copy of the given slice and heapifies it
*/
func Heapify[T any](elements []T, compare func(*T, *T) bool) PriorityQueue[T] {
	size := len(elements)
	container := make([]T, size+1)
	copy(container, elements)
	container[0], container[size] = container[size], container[0]

	for i := size >> 1; i > 0; i-- {
		siftDown(i, container, size, compare)
	}

	return PriorityQueue[T]{
		compare:   compare,
		container: container,
		cap:       size,
		size:      size,
	}
}

/*
Sets the equality comparer which is required for "Remove" and "Contains" methods.
Implements Queuer.SetEqualityComparer
*/
func (pq *PriorityQueue[T]) SetEqualityComparer(equals func(*T, *T) bool) {
	pq.equals = equals
}

/*
Returns the number of elements in the PriorityQueue.
Implements Queuer.Size and Collectioner.Size
*/
func (pq *PriorityQueue[T]) Size() int {
	return pq.size
}

/*
Returns true if the PriorityQueue is empty. Otherwise, false.
Implements Queuer.Empty and Collectioner.Empty
*/
func (pq *PriorityQueue[T]) Empty() bool {
	return pq.size == 0
}

/*
Pushes the given value into the PriorityQueue.
Implements Queuer.Enqueue
*/
func (pq *PriorityQueue[T]) Enqueue(element T) {
	pq.size++
	if pq.cap < pq.size {
		pq.container = append(pq.container, element)
		pq.cap++
	} else {
		pq.container[pq.size] = element
	}

	siftUp(pq.container, pq.size, pq.compare)
}

/*
Removes the top element (with the highest priority) in the PriorityQueue.
Implements Queuer.Dequeue
*/
func (pq *PriorityQueue[T]) Dequeue() {
	if pq.size == 0 {
		panic("Cannot Dequeue. PriorityQueue is empty")
	}

	pq.container[1], pq.container[pq.size] = pq.container[pq.size], pq.container[1]
	pq.size--
	siftDown(1, pq.container, pq.size, pq.compare)
}

/*
Returns a reference to the top element (with the highest priority) without removing it.
Implements Queuer.Peek
*/
func (pq *PriorityQueue[T]) Peek() *T {
	if pq.size == 0 {
		panic("Cannot Peek. PriorityQueue is empty")
	}

	return &pq.container[1]
}

/*
Enqueues the given value into the Priority Queue. Always returns true.
Implements Collectioner.Add
*/
func (pq *PriorityQueue[T]) Add(element T) bool {
	pq.Enqueue(element)
	return true
}

/*
Removes the first occurrence of the given value. Returns true if an element of the same value was found and
removed. If not, returns false. When an element is removed, extra operations are performed to restore the
correct order. Time complexity is O(log n).
Implements Collectioner.Remove
*/
func (pq *PriorityQueue[T]) Remove(element T) bool {
	if pq.equals == nil {
		panic("Cannot Remove. Equality comparer was not set")
	}

	i := 1
	for ; i <= pq.size; i++ {
		if pq.equals(&pq.container[i], &element) {
			break
		}
	}

	if i > pq.size {
		return false
	}

	pq.container[i], pq.container[pq.size] = pq.container[pq.size], pq.container[i]
	pq.size--
	siftDown(i, pq.container, pq.size, pq.compare)

	return true
}

/*
Returns true if an element with the same value as the given value exists. Otherwise, returns false.
Implements Collectioner.Contains
*/
func (pq *PriorityQueue[T]) Contains(element T) bool {
	if pq.equals == nil {
		panic("Cannot execute Contains. Equality comparer was not set")
	}

	for i := 1; i <= pq.size; i++ {
		if pq.equals(&pq.container[i], &element) {
			return true
		}
	}

	return false
}

/*
Empties the PriorityQueue. It does not actually deallocates all the memory that it was using before. It simply
sets the size counter to 0 and overwrites any existing data as new values are enqueued. Therefore, the time
complexity is O(1).
Implements Queuer.Clear and Collectioner.Clear
*/
func (pq *PriorityQueue[T]) Clear() {
	pq.size = 0
}

/*
Executes the given "do" function on each element in the PriorityQueue. After the elements are updated, the
internal array is heapified again in order to restore the appropriate order. Time complexity is O(n).
Implements Queuer.ForEach and Collectioner.ForEach
*/
func (pq *PriorityQueue[T]) ForEach(do func(*T)) {
	for i := 1; i <= pq.size; i++ {
		do(&pq.container[i])
	}

	for i := pq.size >> 1; i > 0; i-- {
		siftDown(i, pq.container, pq.size, pq.compare)
	}
}
