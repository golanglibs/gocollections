package priorityqueue

import (
	"testing"

	"github.com/golanglibs/goassert"
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/queue"
	"github.com/golanglibs/gocollections/testhelpers"
)

func data(val int) testhelpers.MockStruct {
	return testhelpers.MockStruct{Prop: val}
}

func verifyPq(
	t *testing.T,
	expectedOrder []testhelpers.MockStruct,
	actual *PriorityQueue[testhelpers.MockStruct]) {
	t.Helper()

	size := len(expectedOrder)
	i := 0
	for ; i < size && actual.size > 0; i++ {
		goassert.Equal(t, expectedOrder[i], actual.container[1])
		actual.Dequeue()
	}

	goassert.Equal(t, size, i)
	goassert.Equal(t, actual.size, 0)
}

func test_queuer[T any](pq queue.Queuer[T]) {}

func test_collectioner[T any](pq generic.Collectioner[T]) {}

func equals(s0 *testhelpers.MockStruct, s1 *testhelpers.MockStruct) bool {
	return s0.Prop == s1.Prop
}

func compare(s0 *testhelpers.MockStruct, s1 *testhelpers.MockStruct) bool {
	return s0.Prop < s1.Prop
}

func Test_NewShouldCreate_EmptyPriorityQueue(t *testing.T) {
	pq := New(compare)
	goassert.NotNil(t, pq.compare)
	goassert.SliceLength(t, pq.container, 1)
	goassert.Equal(t, 0, pq.cap)
	goassert.Equal(t, 0, pq.size)
}

func Test_HeapifyShouldCorrectlyOrder_GivenArray(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 14},
		{Prop: 16},
		{Prop: 5},
		{Prop: 23},
		{Prop: 7},
		{Prop: 10},
	}
	pq := Heapify(arr, compare)

	correct_order := []testhelpers.MockStruct{
		{Prop: 5},
		{Prop: 7},
		{Prop: 10},
		{Prop: 14},
		{Prop: 16},
		{Prop: 23},
	}

	verifyPq(t, correct_order, &pq)
}

func Test_HeapifyShouldCorrectlyCreatePriorityQueue_GivenEmptyArray(t *testing.T) {
	arr := []testhelpers.MockStruct{}
	pq := Heapify(arr, compare)

	goassert.Equal(t, 0, pq.size)
	goassert.SliceLength(t, pq.container, 1)
}

func Test_SizeShouldReturnCorrectSize(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 14},
		{Prop: 16},
		{Prop: 5},
		{Prop: 23},
		{Prop: 7},
		{Prop: 10},
	}
	pq := Heapify(arr, compare)
	pq.Enqueue(data(21))

	goassert.Equal(t, 7, pq.Size())
	goassert.SliceLength(t, pq.container, 8)
	goassert.Equal(t, 7, pq.cap)
}

func Test_EmptyShouldReturnTrue_GivenEmptyPriorityQueue(t *testing.T) {
	pq := New(compare)

	goassert.True(t, pq.Empty())
}

func Test_EmptyShouldReturnFalse_GivenNonEmptyPriorityQueue(t *testing.T) {
	pq := New(compare)
	pq.Enqueue(data(10))

	goassert.False(t, pq.Empty())
}

func Test_EnqueueShouldAddProvidedValueAtCorrectPosition_GivenEmptyPriorityQueue(t *testing.T) {
	pq := New(compare)
	pq.Enqueue(data(16))
	pq.Enqueue(data(14))
	pq.Enqueue(data(20))

	correct_order := []testhelpers.MockStruct{
		{Prop: 14},
		{Prop: 16},
		{Prop: 20},
	}

	verifyPq(t, correct_order, &pq)
}

func Test_EnqueueShouldAddProvidedValueAtCorrectPosition_GivenNonEmptyPriorityQueue(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 14},
		{Prop: 16},
		{Prop: 5},
		{Prop: 23},
		{Prop: 7},
		{Prop: 10},
	}
	pq := Heapify(arr, compare)
	pq.Enqueue(data(21))

	correct_order := []testhelpers.MockStruct{
		{Prop: 5},
		{Prop: 7},
		{Prop: 10},
		{Prop: 14},
		{Prop: 16},
		{Prop: 21},
		{Prop: 23},
	}

	verifyPq(t, correct_order, &pq)
}

func Test_DequeueShouldRemoveTheCorrectValue(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 14},
		{Prop: 16},
		{Prop: 5},
		{Prop: 23},
		{Prop: 7},
		{Prop: 10},
	}
	pq := Heapify(arr, compare)
	pq.Enqueue(data(21))

	pq.Dequeue()

	correct_order := []testhelpers.MockStruct{
		{Prop: 7},
		{Prop: 10},
		{Prop: 14},
		{Prop: 16},
		{Prop: 21},
		{Prop: 23},
	}

	goassert.Equal(t, 6, pq.size)
	goassert.Equal(t, 7, pq.cap)
	verifyPq(t, correct_order, &pq)
}

func Test_DequeueShouldPanic_GivenEmptyPriorityQueue(t *testing.T) {
	pq := New(compare)
	goassert.PanicWithError(t, "Cannot Dequeue. PriorityQueue is empty", func() {
		pq.Dequeue()
	})
}

func Test_PeekShouldReturnCorrectValue_ButNotRemoveIt(t *testing.T) {
	pq := New(compare)
	pq.Enqueue(data(16))
	pq.Enqueue(data(14))

	goassert.Equal(t, data(14), *pq.Peek())
	goassert.Equal(t, data(14), pq.container[1])
	goassert.Equal(t, 2, pq.size)
}

func Test_PeekShouldPanic_GivenEmptyPriorityQueue(t *testing.T) {
	pq := New(compare)
	goassert.PanicWithError(t, "Cannot Peek. PriorityQueue is empty", func() {
		pq.Peek()
	})
}

func Test_AddShouldEnqueueGivenValueAndReturnTrue(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 14},
		{Prop: 16},
		{Prop: 5},
		{Prop: 23},
		{Prop: 7},
		{Prop: 10},
	}
	pq := Heapify(arr, compare)
	result := pq.Add(data(21))

	correct_order := []testhelpers.MockStruct{
		{Prop: 5},
		{Prop: 7},
		{Prop: 10},
		{Prop: 14},
		{Prop: 16},
		{Prop: 21},
		{Prop: 23},
	}

	verifyPq(t, correct_order, &pq)
	goassert.True(t, result)
}

func Test_RemoveShouldRemoveFirstOccurrenceOfGivenValueAndReturnTrue(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 14},
		{Prop: 16},
		{Prop: 5},
		{Prop: 23},
		{Prop: 7},
		{Prop: 5},
		{Prop: 10},
		{Prop: 5},
	}
	pq := Heapify(arr, compare)
	pq.SetEqualityComparer(equals)
	result := pq.Remove(data(5))

	correct_order := []testhelpers.MockStruct{
		{Prop: 5},
		{Prop: 5},
		{Prop: 7},
		{Prop: 10},
		{Prop: 14},
		{Prop: 16},
		{Prop: 23},
	}

	goassert.True(t, result)
	goassert.Equal(t, 7, pq.size)
	verifyPq(t, correct_order, &pq)
}

func Test_RemoveShouldReturnFalse_IfGivenValueDoesNotExist(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 16},
		{Prop: 14},
		{Prop: 23},
	}
	pq := Heapify(arr, compare)
	pq.SetEqualityComparer(equals)
	result := pq.Remove(data(5))

	goassert.False(t, result)
	goassert.Equal(t, 3, pq.size)
	goassert.Equal(t, 14, pq.container[1].Prop)
}

func Test_RemoveShouldReturnFalse_GivenEmptyPriorityQueue(t *testing.T) {
	pq := New(compare)
	pq.SetEqualityComparer(equals)
	result := pq.Remove(data(5))

	goassert.False(t, result)
	goassert.Equal(t, 0, pq.size)
}

func Test_RemoveShouldPanic_IfEqualityComparerIsNotSet(t *testing.T) {
	pq := New(compare)
	goassert.PanicWithError(t, "Cannot Remove. Equality comparer was not set", func() {
		pq.Remove(data(10))
	})
}

func Test_ContainsShouldReturnTrue_IfGivenValueExists(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 16},
		{Prop: 14},
		{Prop: 23},
	}
	pq := Heapify(arr, compare)
	pq.SetEqualityComparer(equals)

	goassert.True(t, pq.Contains(data(14)))
}

func Test_ContainsShouldReturnFalse_IfGivenValueDoesNotExist(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 16},
		{Prop: 14},
		{Prop: 23},
	}
	pq := Heapify(arr, compare)
	pq.SetEqualityComparer(equals)

	goassert.False(t, pq.Contains(data(20)))
}

func Test_ContainsShouldReturnFalse_IfEmpty(t *testing.T) {
	pq := New(compare)
	pq.SetEqualityComparer(equals)

	goassert.False(t, pq.Contains(data(20)))
}

func Test_ContainsShouldPanic_IfEqualityComparerIsNotSet(t *testing.T) {
	pq := New(compare)
	goassert.PanicWithError(t, "Cannot execute Contains. Equality comparer was not set", func() {
		pq.Contains(data(0))
	})
}

func Test_ClearShouldResetPriorityQueue(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 16},
		{Prop: 14},
		{Prop: 23},
	}
	pq := Heapify(arr, compare)
	pq.Clear()
	pq.Enqueue(data(4))
	pq.Enqueue(data(1))
	pq.Enqueue(data(3))
	pq.Enqueue(data(2))

	correct_order := []testhelpers.MockStruct{
		{Prop: 1},
		{Prop: 2},
		{Prop: 3},
		{Prop: 4},
	}

	goassert.Equal(t, 4, pq.size)
	verifyPq(t, correct_order, &pq)
}

func Test_ForEachShouldExecuteGivenOperationOnEachElementAndCorrectlyReorder(t *testing.T) {
	arr := []testhelpers.MockStruct{
		{Prop: 14},
		{Prop: 16},
		{Prop: 23},
		{Prop: 7},
		{Prop: 5},
		{Prop: 10},
	}
	pq := Heapify(arr, compare)

	pq.ForEach(func(v *testhelpers.MockStruct) {
		v.Prop *= -1
	})

	correct_order := []testhelpers.MockStruct{
		{Prop: -23},
		{Prop: -16},
		{Prop: -14},
		{Prop: -10},
		{Prop: -7},
		{Prop: -5},
	}

	verifyPq(t, correct_order, &pq)
}

func Test_PriorityQueueShouldImplementQueuerInterface(t *testing.T) {
	pq := New(compare)
	test_queuer[testhelpers.MockStruct](&pq)
}

func Test_PriorityQueueShouldImplementCollectionerInterface(t *testing.T) {
	pq := New(compare)
	test_collectioner[testhelpers.MockStruct](&pq)
}
