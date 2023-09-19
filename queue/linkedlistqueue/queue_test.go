package linkedlistqueue

import (
	"testing"

	"github.com/golanglibs/goassert"
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/queue"
	"github.com/golanglibs/gocollections/testhelpers"
)

func testQueuer[T any](q queue.Queuer[T]) {}

func testCollectioner[T any](c generic.Collectioner[T]) {}

func verifyQueue[T any](t *testing.T, expectedElements []T, actual *Queue[T]) {
	t.Helper()

	goassert.Equal(t, len(expectedElements), actual.container.Size())
	for i, e := range expectedElements {
		goassert.DeepEqual(t, e, *actual.container.At(i))
	}
}

func NewShouldCreateEmptyQueue_GivenNoElements(t *testing.T) {
	queue := New[int]()

	goassert.Equal(t, 0, queue.container.Size())
}

func NewShouldCreateQueue_WithGivenElements(t *testing.T) {
	queue := New(10, 16, 14)

	expectedElements := []int{10, 16, 14}
	verifyQueue(t, expectedElements, &queue)
}

func NewOfAnyShouldCreateEmptyQueue_GivenNoElements(t *testing.T) {
	queue := NewOfAny[int]()

	goassert.Equal(t, 0, queue.container.Size())
}

func NewOfAnyShouldCreateQueue_WithGivenElements(t *testing.T) {
	queue := NewOfAny(10, 16, 14)

	expectedElements := []int{10, 16, 14}
	verifyQueue(t, expectedElements, &queue)
}

func NewFromCollectionShouldCreateEmptyQueue_GivenNoElements(t *testing.T) {
	collection := testhelpers.NewMockCollection[int]()
	queue := NewFromCollection[int](collection)

	goassert.Equal(t, 0, queue.container.Size())
}

func NewFromCollectionShouldCreateQueue_WithGivenElements(t *testing.T) {
	collection := testhelpers.NewMockCollection(10, 16, 14)
	queue := NewFromCollection[int](collection)

	expectedElements := []int{10, 16, 14}
	verifyQueue(t, expectedElements, &queue)
}

func NewOfAnyFromCollectionShouldCreateEmptyQueue_GivenNoElements(t *testing.T) {
	collection := testhelpers.NewMockCollection[int]()
	queue := NewOfAnyFromCollection[int](collection)

	goassert.Equal(t, 0, queue.container.Size())
}

func NewOfAnyFromCollectionShouldCreateQueue_WithGivenElements(t *testing.T) {
	collection := testhelpers.NewMockCollection(10, 16, 14)
	queue := NewOfAnyFromCollection[int](collection)

	expectedElements := []int{10, 16, 14}
	verifyQueue(t, expectedElements, &queue)
}

func SetEqualityComparerShouldSetGivenEqualityComparer(t *testing.T) {
	queue := NewOfAny(10, 16, 14)

	queue.SetEqualityComparer(func(a *int, b *int) bool { return a == b })
	goassert.True(t, queue.container.Contains(16))
}

func SizeShouldReturnCorrectLengthOfQueue(t *testing.T) {
	queue := New(10, 16, 14)

	goassert.Equal(t, 3, queue.Size())
}

func EmptyShouldReturnTrue_IfQueueIsEmpty(t *testing.T) {
	queue := New[int]()

	goassert.True(t, queue.Empty())
}

func EmptyShouldReturnFalse_IfQueueIsNotEmpty(t *testing.T) {
	queue := New(10, 16, 14)

	goassert.False(t, queue.Empty())
}

func Test_EnqueueShouldAddGivenValueToBackOfQueue(t *testing.T) {
	queue := New[int]()

	queue.Enqueue(10)
	queue.Enqueue(16)
	queue.Enqueue(14)

	expectedElements := []int{10, 16, 14}
	verifyQueue(t, expectedElements, &queue)
}

func Test_DequeueShouldRemoveFirstEnqueuedElement_IfQueueIsNotEmpty(t *testing.T) {
	queue := New[int]()

	queue.Enqueue(10)
	queue.Enqueue(16)
	queue.Enqueue(14)

	queue.Dequeue()
	goassert.Equal(t, 16, *queue.container.Front())
	goassert.Equal(t, 2, queue.container.Size())
}

func Test_DequeueShouldPanic_IfQueueIsEmpty(t *testing.T) {
	queue := New[int]()

	expectedError := "Queue.Dequeue failed because queue is empty"
	goassert.PanicWithError(t, expectedError, func() { queue.Dequeue() })
}

func Test_PeekShouldReturnElementAtFrontOfQueue_IfQueueIsNotEmpty(t *testing.T) {
	queue := New[int]()

	queue.Enqueue(16)
	queue.Enqueue(5)
	queue.Enqueue(14)
	peeked := queue.Peek()

	goassert.Equal(t, 16, *peeked)
}

func Test_PeekShouldPanic_IfQueueIsEmpty(t *testing.T) {
	queue := New[int]()

	expectedError := "Queue.Peek failed because queue is empty"
	goassert.PanicWithError(t, expectedError, func() { queue.Peek() })
}

func Test_AddShouldReturnTrueAndAddGivenElementToBackOfQueue(t *testing.T) {
	queue := New(10, 5, 16)

	addResult := queue.Add(14)

	expectedElements := []int{10, 5, 16, 14}
	goassert.True(t, addResult)
	verifyQueue(t, expectedElements, &queue)
}

func Test_RemoveShouldReturnTrueRemoveGivenElement_IfElementExistsInQueue(t *testing.T) {
	queue := New(14, 16, 1)

	removeResult := queue.Remove(1)

	expectedElements := []int{14, 16}
	goassert.True(t, removeResult)
	verifyQueue(t, expectedElements, &queue)
}

func Test_RemoveShouldReturnFalse_IfGivenElementDoesNotExist(t *testing.T) {
	queue := New(14, 16, 5)

	removeResult := queue.Remove(1)

	expectedElements := []int{14, 16, 5}
	goassert.False(t, removeResult)
	verifyQueue(t, expectedElements, &queue)
}

func Test_ContainsShouldReturnTrue_IfGivenElementExists(t *testing.T) {
	queue := New(14, 16, 5)

	contains := queue.Contains(16)

	goassert.True(t, contains)
}

func Test_ContainsShouldReturnFalse_IfGivenElementDoesNotExist(t *testing.T) {
	queue := New(14, 16, 5)

	contains := queue.Contains(1)

	goassert.False(t, contains)
}

func Test_ClearShouldEmptyQueue(t *testing.T) {
	queue := New(1, 1, 1)

	queue.Clear()

	goassert.Equal(t, 0, queue.container.Size())
}

func Test_ForEachShouldIterateSequentially_And_ExecuteGivenFunction(t *testing.T) {
	queue := New[string]()

	queue.Enqueue("go")
	queue.Enqueue("is")
	queue.Enqueue("awesome")

	finalString := ""
	queue.ForEach(func(s *string) {
		finalString += *s + " "
	})

	goassert.Equal(t, "go is awesome ", finalString)
}

func Test_LinkedListQueueShouldImplementQueuer(t *testing.T) {
	queue := New[int]()
	testQueuer[int](&queue)
}

func Test_LinkedListQueueShouldImplementCollectioner(t *testing.T) {
	queue := New[int]()
	testCollectioner[int](&queue)
}
