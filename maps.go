package x

// HasKey returns whether k is present in the map m.
func HasKey[M ~map[K]V, K comparable, V any](m M, k K) bool {
	_, ok := m[k]
	return ok
}
