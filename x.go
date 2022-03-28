package x

// Check panics if one of its arguments is a non-nil error.
func Check(a ...any) {
	for _, v := range a {
		if err, ok := v.(error); ok && err != nil {
			panic(err)
		}
	}
}

// Ptr returns a pointer to v.
func Ptr[T any](v T) *T {
	return &v
}

// Reverse reverses a slice.
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
