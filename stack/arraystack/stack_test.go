package arraystack

import (
	"testing"

	"github.com/golanglibs/goassert"
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/stack"
	"github.com/golanglibs/gocollections/testhelpers"
)

func testStacker[T any](s stack.Stacker[T]) {}

func testCollectioner[T any](c generic.Collectioner[T]) {}

func verifyStack[T any](t *testing.T, expectedElements []T, actual *Stack[T]) {
	t.Helper()

	goassert.Equal(t, len(expectedElements), actual.container.Size())
	for i, e := range expectedElements {
		goassert.DeepEqual(t, e, *actual.container.At(i))
	}
}

func NewShouldCreateEmptyStack_GivenNoElements(t *testing.T) {
	stack := New[int]()

	goassert.Equal(t, 0, stack.container.Size())
}

func NewShouldCreateStack_WithGivenElements(t *testing.T) {
	stack := New(10, 16, 14)

	expectedElements := []int{10, 16, 14}
	verifyStack(t, expectedElements, &stack)
}

func NewOfAnyShouldCreateEmptyStack_GivenNoElements(t *testing.T) {
	stack := NewOfAny[int]()

	goassert.Equal(t, 0, stack.container.Size())
}

func NewOfAnyShouldCreateStack_WithGivenElements(t *testing.T) {
	stack := NewOfAny(10, 16, 14)

	expectedElements := []int{10, 16, 14}
	verifyStack(t, expectedElements, &stack)
}

func NewFromCollectionShouldCreateEmptyStack_GivenNoElements(t *testing.T) {
	collection := testhelpers.NewMockCollection[int]()
	stack := NewFromCollection[int](collection)

	goassert.Equal(t, 0, stack.container.Size())
}

func NewFromCollectionShouldCreateStack_WithGivenElements(t *testing.T) {
	collection := testhelpers.NewMockCollection(10, 16, 14)
	stack := NewFromCollection[int](collection)

	expectedElements := []int{10, 16, 14}
	verifyStack(t, expectedElements, &stack)
}

func NewOfAnyFromCollectionShouldCreateEmptyStack_GivenNoElements(t *testing.T) {
	collection := testhelpers.NewMockCollection[int]()
	stack := NewOfAnyFromCollection[int](collection)

	goassert.Equal(t, 0, stack.container.Size())
}

func NewOfAnyFromCollectionShouldCreateStack_WithGivenElements(t *testing.T) {
	collection := testhelpers.NewMockCollection(10, 16, 14)
	stack := NewOfAnyFromCollection[int](collection)

	expectedElements := []int{10, 16, 14}
	verifyStack(t, expectedElements, &stack)
}

func SetEqualityComparerShouldSetGivenEqualityComparer(t *testing.T) {
	stack := NewOfAny(10, 16, 14)

	stack.SetEqualityComparer(func(a *int, b *int) bool { return a == b })
	goassert.True(t, stack.container.Contains(16))
}

func SizeShouldReturnCorrectLengthOfStack(t *testing.T) {
	stack := New(10, 16, 14)

	goassert.Equal(t, 3, stack.Size())
}

func EmptyShouldReturnTrue_IfStackIsEmpty(t *testing.T) {
	stack := New[int]()

	goassert.True(t, stack.Empty())
}

func EmptyShouldReturnFalse_IfStackIsNotEmpty(t *testing.T) {
	stack := New(10, 16, 14)

	goassert.False(t, stack.Empty())
}

func Test_PushShouldAddGivenValueToBackOfStack(t *testing.T) {
	stack := New[int]()

	stack.Push(10)
	stack.Push(16)
	stack.Push(14)

	expectedElements := []int{10, 16, 14}
	verifyStack(t, expectedElements, &stack)
}

func Test_PopShouldRemove_IfStackIsNotEmpty(t *testing.T) {
	stack := New[int]()

	stack.Push(10)
	stack.Push(16)
	stack.Push(14)

	stack.Pop()
	goassert.Equal(t, 16, *stack.container.Back())
	goassert.Equal(t, 2, stack.container.Size())
}

func Test_PopShouldPanic_IfStackIsEmpty(t *testing.T) {
	stack := New[int]()

	expectedError := "Stack.Pop failed because stack is empty"
	goassert.PanicWithError(t, expectedError, func() { stack.Pop() })
}

func Test_PeekShouldReturnElementAtTopOfStack_IfStackIsNotEmpty(t *testing.T) {
	stack := New[int]()

	stack.Push(16)
	stack.Push(5)
	stack.Push(14)
	peeked := stack.Peek()

	goassert.Equal(t, 14, *peeked)
}

func Test_PeekShouldPanic_IfStackIsEmpty(t *testing.T) {
	stack := New[int]()

	expectedError := "Stack.Peek failed because stack is empty"
	goassert.PanicWithError(t, expectedError, func() { stack.Peek() })
}

func Test_AddShouldReturnTrueAndAddGivenElementToTopOfStack(t *testing.T) {
	stack := New(10, 5, 16)

	addResult := stack.Add(14)

	expectedElements := []int{10, 5, 16, 14}
	goassert.True(t, addResult)
	verifyStack(t, expectedElements, &stack)
}

func Test_RemoveShouldReturnTrueRemoveGivenElement_IfElementExistsInStack(t *testing.T) {
	stack := New(14, 16, 1)

	removeResult := stack.Remove(1)

	expectedElements := []int{14, 16}
	goassert.True(t, removeResult)
	verifyStack(t, expectedElements, &stack)
}

func Test_RemoveShouldReturnFalse_IfGivenElementDoesNotExist(t *testing.T) {
	stack := New(14, 16, 5)

	removeResult := stack.Remove(1)

	expectedElements := []int{14, 16, 5}
	goassert.False(t, removeResult)
	verifyStack(t, expectedElements, &stack)
}

func Test_ContainsShouldReturnTrue_IfGivenElementExists(t *testing.T) {
	stack := New(14, 16, 5)

	contains := stack.Contains(16)

	goassert.True(t, contains)
}

func Test_ContainsShouldReturnFalse_IfGivenElementDoesNotExist(t *testing.T) {
	stack := New(14, 16, 5)

	contains := stack.Contains(1)

	goassert.False(t, contains)
}

func Test_ClearShouldEmptyStack(t *testing.T) {
	stack := New(1, 1, 1)

	stack.Clear()

	goassert.Equal(t, 0, stack.container.Size())
}

func Test_ForEachShouldIterateSequentially_And_ExecuteGivenFunction(t *testing.T) {
	stack := New[string]()

	stack.Push("awesome")
	stack.Push("is")
	stack.Push("go")

	finalString := ""
	stack.ForEach(func(s *string) {
		finalString += *s + " "
	})

	goassert.Equal(t, "awesome is go ", finalString)
}

func Test_LinkedListStackShouldImplementStacker(t *testing.T) {
	stack := New[int]()
	testStacker[int](&stack)
}

func Test_LinkedListStackShouldImplementCollectioner(t *testing.T) {
	stack := New[int]()
	testCollectioner[int](&stack)
}
