package hashset

import (
	"testing"

	"github.com/golanglibs/goassert"
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/set"
	"github.com/golanglibs/gocollections/testhelpers"
)

func testSeter[K comparable](s set.Seter[K]) {}

func testCollectioner[K comparable](c generic.Collectioner[K]) {}

func Test_NewShouldCreateEmptySet_GivenNoElements(t *testing.T) {
	set := New[string]()

	goassert.EmptyMap(t, set.container)
}

func Test_NewShouldCreateSet_WithGivenElements(t *testing.T) {
	set := New(1, 5, 3, 5)
	expectedElements := []int{1, 3, 5}

	goassert.MapLength(t, set.container, len(expectedElements))

	for _, v := range expectedElements {
		goassert.MapContainsKey(t, set.container, v)
	}
}

func Test_FromCollectionShouldCreateSet_WithElementsOfGivenCollection(t *testing.T) {
	collection := testhelpers.NewMockCollection(3, 3, 5, 7, 10)
	set := NewFromCollection[int](collection)
	expectedElements := []int{3, 5, 7, 10}

	goassert.MapLength(t, set.container, len(expectedElements))

	for _, v := range expectedElements {
		goassert.MapContainsKey(t, set.container, v)
	}
}

func Test_SizeShouldReturnCorrectNumberOfMembers(t *testing.T) {
	set := New(1, 1, 1, 3, 7, 3)
	goassert.MapLength(t, set.container, 3)
}

func Test_EmptyShouldReturnTrue_GivenEmptySet(t *testing.T) {
	set := New[int]()
	goassert.True(t, set.Empty())
}

func Test_AddShouldAddGivenValueAndReturnTrue_IfGivenValueDoesNotExist(t *testing.T) {
	valueToAdd := 1
	set := New[int]()

	goassert.True(t, set.Add(valueToAdd))
	goassert.MapContainsKey(t, set.container, valueToAdd)
}

func Test_AddShouldNotAddGivenValueAndReturnFalse_IfGivenValueExists(t *testing.T) {
	valueToAdd := 1
	set := New(valueToAdd)

	goassert.False(t, set.Add(valueToAdd))
	goassert.MapLength(t, set.container, 1)
}

func Test_RemoveShouldRemoveGivenValueAndReturnTrue_IfGivenValueExists(t *testing.T) {
	valueToRemove := 1
	set := New(valueToRemove)

	goassert.True(t, set.Remove(valueToRemove))
	goassert.EmptyMap(t, set.container)
}

func Test_RemoveShouldReturnFalse_IfGivenValueDoesNotExist(t *testing.T) {
	valueToRemove := 0
	set := New(1)

	goassert.False(t, set.Remove(valueToRemove))
	goassert.MapLength(t, set.container, 1)
}

func Test_ContainsShouldReturnTrue_IfGivenValueExists(t *testing.T) {
	value := 1
	set := New(value)

	goassert.True(t, set.Contains(value))
}

func Test_ContainsShouldReturnFalse_IfGivenValueDoesNotExist(t *testing.T) {
	value := 1
	set := New(12)

	goassert.False(t, set.Contains(value))
}

func Test_EqualsShouldReturnTrue_IfBothSetsHaveSameMembers(t *testing.T) {
	set0 := New(4, 3, 7)
	set1 := New(7, 3, 4)

	goassert.True(t, set0.Equals(&set1))
}

func Test_EqualsShouldReturnFalse_IfBothSetsHaveDifferentMembers(t *testing.T) {
	set0 := New(1, 3, 7)
	set1 := New(7, 3, 4)

	goassert.False(t, set0.Equals(&set1))
}

func Test_EqualsShouldReturnFalse_IfBothSetsHaveDifferentSize(t *testing.T) {
	set0 := New(1, 3, 7, 4)
	set1 := New(7, 3, 4)

	goassert.False(t, set0.Equals(&set1))
}

func Test_IntersectsShouldReturnTrue_IfSetsHaveCommonMembers(t *testing.T) {
	set0 := New(4, 3, 7, 1)
	set1 := New(7, 3, 4)

	goassert.True(t, set0.Intersects(&set1))
}

func Test_IntersectsShouldReturnFalse_IfSetsHaveNoMembersInCommon(t *testing.T) {
	set0 := New(10, 9, 8, 16)
	set1 := New(7, 3, 4)

	goassert.False(t, set0.Intersects(&set1))
}

func Test_GetIntersectionShouldReturnEmptySet_IfSetsHaveNoMembersInCommon(t *testing.T) {
	set0 := New(10, 9, 8, 16)
	set1 := New(7, 3, 4)

	intersection := set0.GetIntersection(&set1)

	goassert.Equal(t, 0, intersection.Size())
}

func Test_GetIntersectionShouldReturnSetWithCommonMembers_IfSetsHaveCommonMembers(t *testing.T) {
	set0 := New(10, 9, 8, 16)
	set1 := New(10, 16, 4, 0)

	intersection := set0.GetIntersection(&set1)

	expectedSet := New(10, 16)
	goassert.True(t, intersection.Equals(&expectedSet))
}

func Test_GetUnionShouldReturnSetWithAllMembersBetweenTwoSets(t *testing.T) {
	set0 := New(10, 3, 7, 16)
	set1 := New(5, 10, 8, 21)

	union := set0.GetUnion(&set1)

	expectedSet := New(10, 3, 5, 7, 8, 16, 21)
	goassert.True(t, union.Equals(&expectedSet))
}

func Test_IsSupersetOfShouldReturnTrue_IfSet0ContainsAllOfMembersOfSet1(t *testing.T) {
	set0 := New(3, 10, 7, 5)
	set1 := New(5, 10)

	goassert.True(t, set0.IsSupersetOf(&set1))
}

func Test_IsSupersetOfShouldReturnTrue_IfSet0IsIdenticalToSet1(t *testing.T) {
	set0 := New(3, 10, 7, 5)
	set1 := New(3, 10, 7, 5)

	goassert.True(t, set0.IsSupersetOf(&set1))
}

func Test_IsSupersetOfShouldReturnFalse_IfSet0HasLessMembersThanSet1(t *testing.T) {
	set0 := New(3, 10, 7)
	set1 := New(3, 10, 7, 5)

	goassert.False(t, set0.IsSupersetOf(&set1))
}

func Test_IsSupersetOfShouldReturnFalse_IfSet0DoesNotHaveAllMembersOfSet1(t *testing.T) {
	set0 := New(3, 10, 7, 5, 8)
	set1 := New(3, 10, 7, 16)

	goassert.False(t, set0.IsSupersetOf(&set1))
}

func Test_IsSubsetOfShouldReturnTrue_IfSet1HasAllMembersOfSet0(t *testing.T) {
	set0 := New(3, 10, 7)
	set1 := New(3, 10, 7, 16)

	goassert.True(t, set0.IsSubsetOf(&set1))
}

func Test_IsSubsetOfShouldReturnTrue_IfSet1IsIdenticalToSet0(t *testing.T) {
	set0 := New(3, 10, 7)
	set1 := New(3, 10, 7)

	goassert.True(t, set0.IsSubsetOf(&set1))
}

func Test_IsSubsetOfShouldReturnFalse_IfSet1HasLessMembersThanSet0(t *testing.T) {
	set0 := New(3, 10, 7, 16)
	set1 := New(3, 10, 7)

	goassert.False(t, set0.IsSubsetOf(&set1))
}

func Test_IsSubsetOfShouldReturnFalse_IfSet1DoesNotHaveAllMembersOfSet0(t *testing.T) {
	set0 := New(3, 10, 7, 21)
	set1 := New(3, 10, 7, 16)

	goassert.False(t, set0.IsSubsetOf(&set1))
}

func Test_ClearShouldEmptyTheSet(t *testing.T) {
	set := New(1, 3, 10)

	set.Clear()

	goassert.EmptyMap(t, set.container)
}

func Test_ForEachShouldIterateThroughTheSet_And_ExecuteGivenFunctionOnEachElement(t *testing.T) {
	set := New(1, 2, 3)
	sum := 0

	set.ForEach(func(member *int) {
		sum += *member
	})

	goassert.Equal(t, 6, sum)
}

func Test_HashSetShouldImplementSeter(t *testing.T) {
	set := New[int]()
	testSeter[int](&set)
}

func Test_HashSetShouldImplementCollectioner(t *testing.T) {
	set := New[int]()
	testCollectioner[int](&set)
}
