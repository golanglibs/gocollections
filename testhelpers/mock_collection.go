package testhelpers

// implements Collectioner
type MockCollection[T any] struct {
	container []T
}

func NewMockCollection[T any](elements ...T) *MockCollection[T] {
	slice := make([]T, len(elements))
	copy(slice, elements)

	return &MockCollection[T]{
		container: slice,
	}
}

func (c *MockCollection[T]) Size() int {
	return len(c.container)
}

func (c *MockCollection[T]) Empty() bool {
	return len(c.container) == 0
}

func (c *MockCollection[T]) Add(element T) bool {
	c.container = append(c.container, element)
	return true
}

func (c *MockCollection[T]) Remove(element T) bool {
	panic("Not implemented")
}

func (c *MockCollection[T]) Contains(element T) bool {
	panic("Not implemented")
}

func (c *MockCollection[T]) Clear() {
	c.container = make([]T, 0)
}

func (c *MockCollection[T]) ForEach(do func(*T)) {
	for _, v := range c.container {
		do(&v)
	}
}
