package arraylist

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

func Test_NewShouldCreateEmptyList_WithDefaultEquals_GivenNoElements(t *testing.T) {
	list := New[int]()

	goassert.SliceLength(t, list.container, 0)
	goassert.NotNil(t, list.equals)
}

func Test_NewShouldCreateList_WithElementsOfGivenElements_And_DefaultEquals(t *testing.T) {
	list := New(3, 16, 10, 7, 5)

	expectedElements := []int{3, 16, 10, 7, 5}
	goassert.DeepEqual(t, expectedElements, list.container)
	goassert.NotNil(t, list.equals)
}

func Test_NewOfAnyShouldCreateEmptyList_WithNilEquals_GivenNoElements(t *testing.T) {
	list := NewOfAny[testhelpers.MockStruct]()

	goassert.SliceLength(t, list.container, 0)
	goassert.Nil(t, list.equals)
}

func Test_NewOfAnyShouldCreateList_WithGivenElements_And_NilEquals(t *testing.T) {
	list := NewOfAny[testhelpers.MockStruct]()

	goassert.SliceLength(t, list.container, 0)
	goassert.Nil(t, list.equals)
}

func Test_NewFromCollectionShouldCreateList_WithElementsOfGivenCollection_And_DefaultEquals(t *testing.T) {
	collection := testhelpers.NewMockCollection(5, 10, 7, 7)
	list := NewFromCollection[int](collection)

	expectedElements := []int{5, 10, 7, 7}

	goassert.DeepEqual(t, expectedElements, list.container)
	goassert.NotNil(t, list.equals)
}

func Test_NewOfAnyFromCollectionShouldCreateList_WithElementsOfGivenCollection_And_NilEquals(t *testing.T) {
	collection := testhelpers.NewMockCollection([]int{3, 4, 2}, []int{7, 4, 1})
	list := NewOfAnyFromCollection[[]int](collection)

	expectedElements := [][]int{{3, 4, 2}, {7, 4, 1}}

	goassert.DeepEqual(t, expectedElements, list.container)
	goassert.Nil(t, list.equals)
}

func Test_SetEqualityComparerShouldSetGivenEqualsFunction(t *testing.T) {
	list := NewOfAny[[]int]()

	list.SetEqualityComparer(func(a *[]int, b *[]int) bool { return true })

	goassert.NotNil(t, list.equals)
}

func Test_AtShouldReturnElementAtGivenIndex(t *testing.T) {
	list := New(3, 10, 7, 16)

	goassert.Equal(t, 10, *list.At(1))
}

func Test_AtShouldPanic_GivenOutOfRangeIndex(t *testing.T) {
	list := New(3, 10, 7, 16)

	outOfRangeIndex := 5
	expectedError :=
		fmt.Sprintf("List.At could not retrieve element because given index %d is out of range", outOfRangeIndex)

	goassert.PanicWithError(t, expectedError, func() { list.At(outOfRangeIndex) })
}

func Test_SetShouldSetGivenValue_AtGivenIndex(t *testing.T) {
	list := New(3, 10, 7, 16)

	list.Set(2, 5)
	goassert.Equal(t, 5, list.container[2])
}

func Test_SetShouldPanic_GivenOutOfRangeIndex(t *testing.T) {
	list := New(3, 10, 7, 16)

	outOfRangeIndex := 10
	expectedError :=
		fmt.Sprintf("List.Set could not set given value because given index %d is out of range", outOfRangeIndex)

	goassert.PanicWithError(t, expectedError, func() { list.Set(outOfRangeIndex, 8) })
}

func Test_SizeShouldReturnTheCorrectLengthOfList(t *testing.T) {
	list := New(10, 10, 10, 10, 10)

	goassert.Equal(t, 5, list.Size())
}

func Test_EmptyShouldReturnTrue_GivenEmptyList(t *testing.T) {
	list := New[int]()
	goassert.True(t, list.Empty())
}

func Test_EmptyShouldReturnFalse_GivenNonEmptyList(t *testing.T) {
	list := New(1, 2, 3)
	goassert.False(t, list.Empty())
}

func Test_FrontShouldReturnRefToFirstElement(t *testing.T) {
	list := New(14, 16)
	goassert.Equal(t, 14, *list.Front())
}

func Test_FrontShouldPanic_GivenEmptyList(t *testing.T) {
	list := New[int]()
	goassert.PanicWithError(t, "ArrayList.Front failed because the list is empty", func() {
		list.Front()
	})
}

func Test_BackshouldReturnRefToLastElement(t *testing.T) {
	list := New(14, 16)
	goassert.Equal(t, 16, *list.Back())
}

func Test_BackShouldPanic_GivenEmptyList(t *testing.T) {
	list := New[int]()
	goassert.PanicWithError(t, "ArrayList.Back failed because the list is empty", func() {
		list.Back()
	})
}

func Test_AddShouldAddGivenValue_And_ReturnTrue(t *testing.T) {
	valueToAdd := 5
	list := New(10, 10)
	addResult := list.Add(valueToAdd)

	expectedElements := []int{10, 10, 5}

	goassert.True(t, addResult)
	goassert.DeepEqual(t, expectedElements, list.container)
}

func Test_RemoveBackShouldRemoveLastElement(t *testing.T) {
	list := New(14, 16, 1)
	list.RemoveBack()

	goassert.Equal(t, 2, list.size)
	goassert.Equal(t, 14, list.container[0])
	goassert.Equal(t, 16, list.container[1])
}

func Test_RemoveShouldPanic_GivenEmptyList(t *testing.T) {
	list := New[int]()
	goassert.PanicWithError(t, "ArrayList.RemoveBack failed because the list is empty", func() {
		list.RemoveBack()
	})
}

func Test_InsertShouldInsertGivenElementAtGivenIndex_And_ReturnTrue_GivenValidIndex(t *testing.T) {
	list := New(10, 10, 3, 7, 16)
	insertIndex := 3
	valueToInsert := 5

	insertResult := list.Insert(insertIndex, valueToInsert)

	expectedElements := []int{10, 10, 3, valueToInsert, 7, 16}
	goassert.True(t, insertResult)
	goassert.DeepEqual(t, expectedElements, list.container)
}

func Test_InsertShouldInsertGivenElementAtBeginningOfList_And_ReturnTrue_GivenIndex0(t *testing.T) {
	list := New(10, 5, 16)
	insertIndex := 0
	valueToInsert := 7

	insertResult := list.Insert(insertIndex, valueToInsert)

	expectedElements := []int{valueToInsert, 10, 5, 16}
	goassert.True(t, insertResult)
	goassert.DeepEqual(t, expectedElements, list.container)
}

func Test_InsertShouldInsertGivenElementAtEndOfList_And_ReturnTrue_GivenMaxIndex(t *testing.T) {
	list := New(10, 5, 16)
	insertIndex := 3
	valueToInsert := 7

	insertResult := list.Insert(insertIndex, valueToInsert)

	expectedElements := []int{10, 5, 16, valueToInsert}
	goassert.True(t, insertResult)
	goassert.DeepEqual(t, expectedElements, list.container)
}

func Test_AddToFrontShouldInsertGivenElementToFrontOfList(t *testing.T) {
	list := New(16)
	list.AddToFront(14)

	expectedElements := []int{14, 16}
	goassert.DeepEqual(t, expectedElements, list.container)
	goassert.Equal(t, 2, list.size)
}

func Test_RemoveFrontShouldRemoveFirstElementOfList(t *testing.T) {
	list := New(1, 14, 16)
	list.RemoveFront()

	expectedElements := []int{14, 16}
	goassert.DeepEqual(t, expectedElements, list.container)
	goassert.Equal(t, 2, list.size)
}

func Test_RemoveFrontShouldPanic_GivenEmptyList(t *testing.T) {
	list := New[int]()
	goassert.PanicWithError(
		t,
		"ArrayList.RemoveAt cannot remove element at index 0 because it is out of range",
		func() {
			list.RemoveFront()
		},
	)
}

func Test_RemoveShouldRemoveGivenElement_And_ReturnTrue_IfGivenElementExistsAtBeginningOfList(t *testing.T) {
	list := New(10, 16, 5)

	removalResult := list.Remove(10)

	expectedElements := []int{16, 5}
	goassert.True(t, removalResult)
	goassert.DeepEqual(t, expectedElements, list.container)
}

func Test_RemoveShouldRemoveGivenElement_And_ReturnTrue_IfGivenElementExistsAtMiddleOfList(t *testing.T) {
	list := New(10, 16, 5)

	removalResult := list.Remove(16)

	expectedElements := []int{10, 5}
	goassert.True(t, removalResult)
	goassert.DeepEqual(t, expectedElements, list.container)
}

func Test_RemoveShouldRemoveGivenElement_And_ReturnTrue_IfGivenElementExistsAtEndOfList(t *testing.T) {
	list := New(10, 16, 5)

	removalResult := list.Remove(5)

	expectedElements := []int{10, 16}
	goassert.True(t, removalResult)
	goassert.DeepEqual(t, expectedElements, list.container)
}

func Test_RemoveShouldRemoveGivenElement_And_ReturnFalse_IfGivenElementDoesNotExist(t *testing.T) {
	list := New(10, 16, 5)

	removalResult := list.Remove(7)

	expectedElements := []int{10, 16, 5}
	goassert.False(t, removalResult)
	goassert.DeepEqual(t, expectedElements, list.container)
}

func Test_RemoveShouldPanic_IfEqualityComparerIsNotSet(t *testing.T) {
	list := NewOfAny(testhelpers.NewMockStruct(10))

	goassert.PanicWithError(t, missingEqualityComparerError, func() {
		list.Remove(testhelpers.NewMockStruct(10))
	})
}

func Test_RemoveAtShouldRemoveElementAtGivenIndex_GivenBeginningIndex(t *testing.T) {
	list := New(10, 16, 5)

	list.RemoveAt(0)

	expectedElements := []int{16, 5}
	goassert.DeepEqual(t, expectedElements, list.container)
	goassert.Equal(t, 2, list.size)
}

func Test_RemoveAtShouldRemoveElementAtGivenIndex_GivenIndexInMiddleOfList(t *testing.T) {
	list := New(10, 16, 5)

	list.RemoveAt(1)

	expectedElements := []int{10, 5}
	goassert.DeepEqual(t, expectedElements, list.container)
	goassert.Equal(t, 2, list.size)
}

func Test_RemoveAtShouldRemoveElementAtGivenIndex_GivenEndIndex(t *testing.T) {
	list := New(10, 16, 5)

	list.RemoveAt(2)

	expectedElements := []int{10, 16}
	goassert.DeepEqual(t, expectedElements, list.container)
	goassert.Equal(t, 2, list.size)
}

func Test_RemoveAtShouldPanicGivenOutOfRangeIndex(t *testing.T) {
	list := New(10, 16, 5)
	goassert.PanicWithError(
		t,
		"ArrayList.RemoveAt cannot remove element at index 10 because it is out of range",
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

func Test_SubListShouldReturnNewCopiedSubList_GivenValidRange(t *testing.T) {
	list := New(16, 10, 5, 16, 10, 5)

	subList := list.SubList(1, 4)
	subArrayList, _ := subList.(*List[int])
	// mutate original list to check if sub list is actually a copy
	list.Set(2, 14)

	expectedList := []int{10, 5, 16}
	goassert.DeepEqual(t, expectedList, subArrayList.container)
}

func Test_SubListShouldCopyWholeList_GivenStartIndexAndLengthAsLastIndex(t *testing.T) {
	list := New(16, 10, 5, 16, 10, 5)

	subList := list.SubList(0, 6)
	subArrayList, _ := subList.(*List[int])
	// mutate original list to check if sub list is actually a copy
	list.Set(2, 14)

	expectedList := []int{16, 10, 5, 16, 10, 5}
	goassert.DeepEqual(t, expectedList, subArrayList.container)
}

func Test_SubListShouldPanic_GivenInvalidRange(t *testing.T) {
	list := New(16, 14, 5, 10)

	expectedError := "List.SubList Cannot create a sub list because invalid range was given"
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

func Test_ClearShouldEmptyCurrentList(t *testing.T) {
	list := New(16, 14, 10)

	list.Clear()

	goassert.Equal(t, 0, list.size)
}

func Test_ForEachShouldIterateThroughTheList_And_ExecuteGivenFunctionOnEachElement(t *testing.T) {
	list := New(10, 16, 5)

	sum := 0
	list.ForEach(func(element *int) {
		sum += *element
	})

	goassert.Equal(t, 31, sum)
}

func Test_ArrayListShouldImplementLister(t *testing.T) {
	list := New[int]()
	testLister[int](&list)
}

func Test_ArrayListShouldImplementCollectioner(t *testing.T) {
	list := New[int]()
	testCollectioner[int](&list)
}
