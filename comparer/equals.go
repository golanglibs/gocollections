package comparer

func DefaultEquals[K comparable](a *K, b *K) bool {
	return *a == *b
}
