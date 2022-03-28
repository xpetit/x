# X

A collection of functions to write concise code.

```go
// Check panics if one of its arguments is a non-nil error.
func Check(a ...any)

// Ptr returns a pointer to v.
func Ptr[T any](v T) *T

// Reverse reverses a slice.
func Reverse[S ~[]E, E any](s S)
```
