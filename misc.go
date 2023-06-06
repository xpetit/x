package x

import (
	"math"
	"sync"
)

// Ptr returns a pointer to v.
func Ptr[T any](v T) *T {
	return &v
}

// Goroutines spawns nb goroutines executing fn. It returns a function that waits for them to finish.
func Goroutines(nb int, fn func()) (wait func()) {
	var wg sync.WaitGroup
	for i := 0; i < nb; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	return wg.Wait
}

// IsNaN32 is the float32 variant of math.IsNaN.
func IsNaN32(f float32) bool { return f != f }

// Less implements "<" for some core types
func Less[T comparable](a, b T) bool {
	switch a := any(a).(type) {

	case int8:
		return a < any(b).(int8)
	case int16:
		return a < any(b).(int16)
	case int32:
		return a < any(b).(int32)
	case int64:
		return a < any(b).(int64)
	case int:
		return a < any(b).(int)

	case uint8:
		return a < any(b).(uint8)
	case uint16:
		return a < any(b).(uint16)
	case uint32:
		return a < any(b).(uint32)
	case uint64:
		return a < any(b).(uint64)
	case uint:
		return a < any(b).(uint)

	case uintptr:
		return a < any(b).(uintptr)

	case string:
		return a < any(b).(string)

	case float32:
		b := any(b).(float32)
		return a < b || (IsNaN32(a) && !IsNaN32(b))
	case float64:
		b := any(b).(float64)
		return a < b || (math.IsNaN(a) && !math.IsNaN(b))
	}

	return false
}
