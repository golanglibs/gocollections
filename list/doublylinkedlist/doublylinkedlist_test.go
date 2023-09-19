package doublylinkedlist

import (
	"fmt"
	"testing"

	"github.com/golanglibs/goassert"
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/list"
	"github.com/golanglibs/gocollections/testhelpers"
)

var missingEqualityComparerError string = "Cannot compute equality of elements since equality comparer is not set"

func testLister[T any](l list.Lister[T]) {}

func testCollectioner[T any](c generic.Collectioner[T]) {}

func verifyDoublyLinkedList[T any](t *testing.T, expected []T, actual *DoublyLinkedList[T]) {
	t.Helper()

	current := actual.head.Next
	for _, v := range expected {
		goassert.DeepEqual(t, v, current.Value)
		current = current.Next
	}
	goassert.Equal(t, actual.tail, current)

	current = actual.tail.Prev
	for i := len(expected) - 1; i >= 0; i-- {
		goassert.DeepEqual(t, expected[i], current.Value)
		current = current.Prev
	}
	goassert.Equal(t, actual.head, current)
}

func Test_NewShouldCreateDoublyLinkedList_WithGivenElements(t *testing.T) {
	list := New(10, 16, 5)

	expectedElements := []int{10, 16, 5}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_NewShouldCreateEmptyDoublyLinkedList_GivenNoElements(t *testing.T) {
	list := New[int]()

	expectedElements := make([]int, 0)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_NewOfAnyShouldCreateDoublyLinkedList_WithGivenElements(t *testing.T) {
	list := NewOfAny(10, 16, 5)

	expectedElements := []int{10, 16, 5}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_NewOfAnyShouldCreateEmptyDoublyLinkedList_GivenNoElements(t *testing.T) {
	list := NewOfAny[string]()

	expectedElements := make([]string, 0)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_NewFromCollectionShouldCreateDoublyLinkedList_WithGivenElements(t *testing.T) {
	collection := testhelpers.NewMockCollection(10, 16, 5)
	list := NewFromCollection[int](collection)

	expectedElements := []int{10, 16, 5}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_NewOfAnyFromCollectionShouldCreateDoublyLinkedList_WithGivenElements(t *testing.T) {
	collection := testhelpers.NewMockCollection([]int{16}, []int{10})
	list := NewOfAnyFromCollection[[]int](collection)

	expectedElements := [][]int{{16}, {10}}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_SetEqualityComparerShouldSetGivenComparer(t *testing.T) {
	list := NewOfAny[int]()

	list.SetEqualityComparer(func(a *int, b *int) bool { return a == b })

	goassert.NotNil(t, list.equals)
}

func Test_AtShouldReturnValue_AtGivenIndex(t *testing.T) {
	list := New(3, 10, 7, 16)

	goassert.Equal(t, 10, *list.At(1))
}

func Test_AtShouldPanic_GivenOutOfRangeIndex(t *testing.T) {
	list := New(3, 10, 7, 16)

	outOfRangeIndex := 5
	expectedError :=
		fmt.Sprintf("DoublyLinkedList.At could not retrieve element because given index %d is out of range", outOfRangeIndex)

	goassert.PanicWithError(t, expectedError, func() { list.At(outOfRangeIndex) })
}

func Test_FrontShouldReturnHeadValue_IfDoublyLinkedListIsNotEmpty(t *testing.T) {
	list := New(16, 10)

	headValue := list.Front()
	goassert.Equal(t, 16, *headValue)
}

func Test_FrontShouldPanic_IfDoublyLinkedListIsEmpty(t *testing.T) {
	list := New[int]()

	expectedError := "DoublyLinkedList.Front failed because the list is empty"
	goassert.PanicWithError(t, expectedError, func() { list.Front() })
}

func Test_TailShouldReturnHeadValue_IfDoublyLinkedListIsNotEmpty(t *testing.T) {
	list := New(16, 10)

	headValue := list.Back()
	goassert.Equal(t, 10, *headValue)
}

func Test_BackShouldPanic_IfDoublyLinkedListIsEmpty(t *testing.T) {
	list := New[int]()

	expectedError := "DoublyLinkedList.Back failed because the list is empty"
	goassert.PanicWithError(t, expectedError, func() { list.Back() })
}

func Test_SetShouldSetGivenValueAtGivenIndex(t *testing.T) {
	list := New(3, 10, 7, 16)

	list.Set(2, 5)
	goassert.Equal(t, 5, list.head.Next.Next.Next.Value)
}

func Test_SetShouldPanic_GivenOutOfRangeIndex(t *testing.T) {
	list := New(3, 10, 7, 16)

	outOfRangeIndex := 10
	expectedError :=
		fmt.Sprintf("DoublyLinkedList.Set could not set given value because given index %d is out of range", outOfRangeIndex)

	goassert.PanicWithError(t, expectedError, func() { list.Set(outOfRangeIndex, 8) })
}

func Test_SizeShouldReturnTheCorrectLengthOfList(t *testing.T) {
	list := New(10, 10, 10, 10, 10)

	goassert.Equal(t, 5, list.Size())
}

func Test_AddShouldAddGivenValue_And_ReturnTrue(t *testing.T) {
	valueToAdd := 5
	list := New(10, 10)
	addResult := list.Add(valueToAdd)

	expectedElements := []int{10, 10, 5}
	goassert.True(t, addResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_AddShouldSetHeadAndTailCorrectly_WhenAddingToEmptyDoublyLinkedList(t *testing.T) {
	valueToAdd := 16
	list := New[int]()
	addResult := list.Add(valueToAdd)

	expectedElements := []int{16}
	goassert.True(t, addResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_AddToFrontShouldAddGivenElement_AheadOfHead(t *testing.T) {
	valueToAdd := 16
	list := New(10, 5)
	list.AddToFront(valueToAdd)

	expectedElements := []int{16, 10, 5}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_AddToFrontShouldSetHeadAndTailCorrectly_WhenAddingToEmptyDoublyLinkedList(t *testing.T) {
	valueToAdd := 16
	list := New[int]()
	list.AddToFront(valueToAdd)

	expectedElements := []int{16}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_InsertShouldInsertGivenElementAtGivenIndex_And_ReturnTrue_GivenValidIndex(t *testing.T) {
	list := New(10, 10, 3, 7, 16)
	insertIndex := 3
	valueToInsert := 5

	insertResult := list.Insert(insertIndex, valueToInsert)

	expectedElements := []int{10, 10, 3, valueToInsert, 7, 16}
	goassert.True(t, insertResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_InsertShouldInsertGivenElementAtBeginningOfList_And_ReturnTrue_GivenIndex0(t *testing.T) {
	list := New(10, 5, 16)
	insertIndex := 0
	valueToInsert := 7

	insertResult := list.Insert(insertIndex, valueToInsert)

	expectedElements := []int{valueToInsert, 10, 5, 16}
	goassert.True(t, insertResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_InsertShouldInsertGivenElementAtEndOfList_And_ReturnTrue_GivenMaxIndex(t *testing.T) {
	list := New(10, 5, 16)
	insertIndex := 3
	valueToInsert := 7

	insertResult := list.Insert(insertIndex, valueToInsert)

	expectedElements := []int{10, 5, 16, valueToInsert}
	goassert.True(t, insertResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_InsertShouldSetHeadAndTailCorrectly_WhenInsertingIntoEmptyDoublyLinkedList(t *testing.T) {
	list := New[int]()
	insertIndex := 0
	valueToInsert := 16

	insertResult := list.Insert(insertIndex, valueToInsert)

	expectedElements := []int{16}
	goassert.True(t, insertResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveShouldRemoveGivenElement_And_ReturnTrue_IfGivenElementExistsAtBeginningOfList(t *testing.T) {
	list := New(10, 16, 5)

	removalResult := list.Remove(10)

	expectedElements := []int{16, 5}
	goassert.True(t, removalResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveShouldRemoveGivenElement_And_ReturnTrue_IfGivenElementExistsAtMiddleOfList(t *testing.T) {
	list := New(10, 16, 5)

	removalResult := list.Remove(16)

	expectedElements := []int{10, 5}
	goassert.True(t, removalResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveShouldRemoveGivenElement_And_ReturnTrue_IfGivenElementExistsAtEndOfList(t *testing.T) {
	list := New(10, 16, 5)

	removalResult := list.Remove(5)

	expectedElements := []int{10, 16}
	goassert.True(t, removalResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveShouldCorrectlySetHeadAndTail_GivenDoublyLinkedListWithTwoElements(t *testing.T) {
	list := New(10, 16)

	removalResult := list.Remove(10)

	expectedElements := []int{16}
	goassert.True(t, removalResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveShouldCorrectlySetHeadAndTail_GivenDoublyLinkedListWithOneElement(t *testing.T) {
	list := New(16)

	removalResult := list.Remove(16)

	goassert.True(t, removalResult)
	goassert.Equal(t, list.tail, list.head.Next)
	goassert.Equal(t, list.head, list.tail.Prev)
	goassert.Equal(t, 0, list.size)
}

func Test_RemoveShouldRemoveGivenElement_And_ReturnFalse_IfGivenElementDoesNotExist(t *testing.T) {
	list := New(10, 16, 5)

	removalResult := list.Remove(7)

	expectedElements := []int{10, 16, 5}
	goassert.False(t, removalResult)
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveShouldPanic_IfEqualityComparerIsNotSet(t *testing.T) {
	list := NewOfAny(testhelpers.NewMockStruct(10))

	goassert.PanicWithError(t, missingEqualityComparerError, func() {
		list.Remove(testhelpers.NewMockStruct(10))
	})
}

func Test_RemoveHeadShouldCorrectlyRemove_GivenListWithThreeOrMoreElements(t *testing.T) {
	list := New(16, 10, 5, 16, 10)

	list.RemoveFront()

	expectedElements := []int{10, 5, 16, 10}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveHeadShouldCorrectlySetHeadAndTail_GivenListWithTwoElements(t *testing.T) {
	list := New(16, 10)

	list.RemoveFront()

	expectedElements := []int{10}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveHeadShouldCorrectlySetHeadAndTail_GivenListWithOneElement(t *testing.T) {
	list := New(16)

	list.RemoveFront()

	goassert.Equal(t, list.tail, list.head.Next)
	goassert.Equal(t, list.head, list.tail.Prev)
	goassert.Equal(t, 0, list.size)
}

func Test_RemoveHeadShouldPanic_GivenEmptyList(t *testing.T) {
	list := New[int]()

	expectedError := "DoublyLinkedList.RemoveFront cannot remove head because list is empty"
	goassert.PanicWithError(t, expectedError, func() { list.RemoveFront() })
}

func Test_RemoveTailShouldCorrectlyRemoveAndReturnTailValue_GivenListWithThreeOrMoreElements(t *testing.T) {
	list := New(16, 10, 5, 16, 10)

	list.RemoveBack()

	expectedElements := []int{16, 10, 5, 16}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveTailShouldCorrectlySetHeadAndTail_GivenListWithTwoElements(t *testing.T) {
	list := New(16, 10)

	list.RemoveBack()

	expectedElements := []int{16}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveTailShouldCorrectlySetHeadAndTail_GivenListWithOneElement(t *testing.T) {
	list := New(16)

	list.RemoveBack()

	goassert.Equal(t, list.tail, list.head.Next)
	goassert.Equal(t, list.head, list.tail.Prev)
	goassert.Equal(t, 0, list.size)
}

func Test_RemoveBackShouldPanic_GivenEmptyList(t *testing.T) {
	list := New[int]()

	expectedError := "DoublyLinkedList.RemoveBack cannot remove tail because list is empty"
	goassert.PanicWithError(t, expectedError, func() { list.RemoveBack() })
}

func Test_RemoveAtShouldRemoveElementAtGivenIndex_GivenBeginningIndex(t *testing.T) {
	list := New(10, 16, 5)

	list.RemoveAt(0)

	expectedElements := []int{16, 5}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveAtShouldRemoveElementAtGivenIndex_GivenIndexInMiddleOfList(t *testing.T) {
	list := New(10, 16, 5)

	list.RemoveAt(1)

	expectedElements := []int{10, 5}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveAtShouldRemoveElementAtGivenIndex_GivenEndIndex(t *testing.T) {
	list := New(10, 5)

	list.RemoveAt(1)

	expectedElements := []int{10}
	verifyDoublyLinkedList(t, expectedElements, &list)
}

func Test_RemoveAtShouldRemoveTheOnlyElement_GivenDoublyLinkedListWithOneElement(t *testing.T) {
	list := New(10)

	list.RemoveAt(0)

	goassert.Equal(t, list.tail, list.head.Next)
	goassert.Equal(t, list.head, list.tail.Prev)
	goassert.Equal(t, 0, list.size)
}

func Test_RemoveAtPanic_GivenOutOfRangeIndex(t *testing.T) {
	list := New(10, 16, 5)
	goassert.PanicWithError(
		t,
		"DoublyLinkedList.RemoveAt cannot remove element at index 10 because it is out of range",
		func() {
			list.RemoveAt(10)
		},
	)
}

func Test_IndexOfShouldReturnElementIndex_IfGivenElementIsFound(t *testing.T) {
	list := New(10, 16, 5)

	elementIndex := list.IndexOf(16)

	goassert.Equal(t, 1, elementIndex)
}

func Test_IndexOfShouldReturnNegativeOne_IfGivenElementIsNotFound(t *testing.T) {
	list := New(10, 16, 5)

	elementIndex := list.IndexOf(7)

	goassert.Equal(t, -1, elementIndex)
}

func Test_ContainsShouldReturnTrue_IfGivenElementIsFound(t *testing.T) {
	list := New(10, 16, 5)

	contains := list.Contains(5)

	goassert.True(t, contains)
}

func Test_ContainsShouldReturnFalse_IfGivenElementIsNotFound(t *testing.T) {
	list := New(10, 16, 5)

	contains := list.Contains(7)

	goassert.False(t, contains)
}

func Test_ClearShouldEmptyCurrentDoublyLinkedList(t *testing.T) {
	list := New(10, 16, 5, 10, 16)

	list.Clear()

	goassert.Equal(t, list.tail, list.head.Next)
	goassert.Equal(t, list.head, list.tail.Prev)
	goassert.Equal(t, 0, list.size)
}

func Test_SubListShouldReturnNewCopiedSubList_GivenValidRange(t *testing.T) {
	list := New(16, 10, 5, 16, 10, 5)

	subList := list.SubList(1, 4)
	subDoublyLinkedList, _ := subList.(*DoublyLinkedList[int])
	// mutate original list to check if sub list is actually a copy
	list.Set(2, 14)

	expectedList := []int{10, 5, 16}
	verifyDoublyLinkedList(t, expectedList, subDoublyLinkedList)
}

func Test_SubListShouldCopyWholeList_GivenStartIndexAndLengthAsLastIndex(t *testing.T) {
	list := New(16, 10, 5, 16, 10, 5)

	subList := list.SubList(0, 6)
	subDoublyLinkedList, _ := subList.(*DoublyLinkedList[int])
	// mutate original list to check if sub list is actually a copy
	list.Set(2, 14)

	expectedList := []int{16, 10, 5, 16, 10, 5}
	verifyDoublyLinkedList(t, expectedList, subDoublyLinkedList)
}

func Test_SubListShouldPanic_GivenInvalidRange(t *testing.T) {
	list := New(16, 14, 5, 10)

	expectedError := "DoublyLinkedList.SubList Cannot create a sub list because invalid range was given"
	testCases := []struct {
		start int
		end   int
	}{
		{-1, 1},
		{1, 5},
		{-1, 10},
		{1, 0},
	}

	for _, c := range testCases {
		goassert.PanicWithError(t, expectedError, func() { list.SubList(c.start, c.end) })
	}
}

func Test_ForEachShouldIterateThroughTheList_And_ExecuteGivenFunctionOnEachElement(t *testing.T) {
	list := New(10, 16, 5)

	sum := 0
	list.ForEach(func(element *int) {
		sum += *element
	})

	goassert.Equal(t, 31, sum)
}

func Test_DoublyLinkedListShouldImplementLister(t *testing.T) {
	list := New[int]()
	testLister[int](&list)
}

func Test_DoublyLinkedListShouldImplementCollectioner(t *testing.T) {
	list := New[int]()
	testCollectioner[int](&list)
}
