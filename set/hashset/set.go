package hashset

import (
	"github.com/golanglibs/gocollections/generic"
	"github.com/golanglibs/gocollections/set"
)

var placeholder interface{} = nil

/*
A hash set that uses native go map under the hood. It implements Seter and Collectioner.
Set is not thread safe
*/
type Set[K comparable] struct {
	container map[K]interface{}
}

/*
Creates an instance of Set with given elements and returns pointer to the instance.
If no elements are given, an empty set is created.
Elements must be comparable
*/
func New[K comparable](elements ...K) Set[K] {
	size := len(elements)
	container := make(map[K]interface{}, size)

	if size > 0 {
		for _, element := range elements {
			container[element] = placeholder
		}
	}

	return Set[K]{
		container: container,
	}
}

/*
Creates an instance of Set with given collection and returns pointer to the instance.
Elements must be comparable
*/
func NewFromCollection[K comparable](c generic.Collectioner[K]) Set[K] {
	size := c.Size()
	container := make(map[K]interface{}, size)

	if size > 0 {
		c.ForEach(func(element *K) {
			container[*element] = placeholder
		})
	}

	return Set[K]{
		container: container,
	}
}

/*
Gets the length of the Set.
Implements Seter.Size and Collectioner.Size
*/
func (s *Set[K]) Size() int {
	return len(s.container)
}

/*
Returns true if the Set is empty.
Implements Seter.Empty and Collectioner.Empty
*/
func (s *Set[K]) Empty() bool {
	return len(s.container) == 0
}

/*
Adds the given element to the Set.
If the element does not exist in the Set, Add will add the given element and return true.
If the element already exists in the Set, Add will not add the given element and return false.
Implements Seter.Add and Collectioner.Add
*/
func (s *Set[K]) Add(element K) bool {
	if _, hasElement := s.container[element]; hasElement {
		return false
	}

	s.container[element] = placeholder
	return true
}

/*
Removes the given element from the Set.
If the given element is found, Remove will delete the element from the Set and return true.
If the given element to remove is not found in the Set, then Remove will return false.
Implements Seter.Remove and Collectioner.Remove
*/
func (s *Set[K]) Remove(element K) bool {
	if _, hasElement := s.container[element]; !hasElement {
		return false
	}

	delete(s.container, element)
	return true
}

/*
Returns true when the given element exists in the Set.
Implements Seter.Contains and Collectioner.Contains
*/
func (s *Set[K]) Contains(element K) bool {
	if _, hasElement := s.container[element]; hasElement {
		return true
	}

	return false
}

/*
Returns true when the given Set has the equal members as the current Set
Implements Seter.Equals
*/
func (s *Set[K]) Equals(set set.Seter[K]) bool {
	if s.Size() != set.Size() {
		return false
	}

	for k := range s.container {
		if !set.Contains(k) {
			return false
		}
	}

	return true
}

/*
Returns true when the given Set has common members with the current Set.
Implements Seter.Intersects
*/
func (s *Set[K]) Intersects(set set.Seter[K]) bool {
	for k := range s.container {
		if set.Contains(k) {
			return true
		}
	}

	return false
}

/*
Returns a new instance of Set with the common members between the current Set and the given Set.
Implements Seter.GetIntersection
*/
func (s *Set[K]) GetIntersection(set set.Seter[K]) set.Seter[K] {
	intersection := &Set[K]{
		container: make(map[K]interface{}),
	}

	for k := range s.container {
		if set.Contains(k) {
			intersection.Add(k)
		}
	}

	return intersection
}

/*
Returns a new instance of Set with all the members of both the current Set and the given Set.
Implements Seter.GetUnion
*/
func (s *Set[K]) GetUnion(set set.Seter[K]) set.Seter[K] {
	union := &Set[K]{
		container: make(map[K]interface{}),
	}

	for k := range s.container {
		union.Add(k)
	}

	set.ForEach(func(member *K) {
		union.Add(*member)
	})

	return union
}

/*
Returns true if the current Set contains all the members of the given Set.
Implements Seter.IsSupersetOf
*/
func (s *Set[K]) IsSupersetOf(set set.Seter[K]) bool {
	if s.Size() < set.Size() {
		return false
	}

	isSuperset := true
	set.ForEach(func(member *K) {
		if !s.Contains(*member) {
			isSuperset = false
		}
	})

	return isSuperset
}

/*
Returns true if the given Set has all the members of the current Set.
Implements Seter.IsSubsetOf
*/
func (s *Set[K]) IsSubsetOf(set set.Seter[K]) bool {
	if s.Size() > set.Size() {
		return false
	}

	for k := range s.container {
		if !set.Contains(k) {
			return false
		}
	}

	return true
}

/*
Clears the current Set so it becomes empty. Under the hood, a new instance of builtin map is assigned as the
new internal container
Implements Seter.Clear
*/
func (s *Set[K]) Clear() {
	s.container = make(map[K]interface{})
}

/*
Iterates through each element in the Set and executes the given "do" function on each element.
Implements Seter.ForEach and Collectioner.ForEach
*/
func (s *Set[K]) ForEach(do func(*K)) {
	for k := range s.container {
		do(&k)
	}
}
