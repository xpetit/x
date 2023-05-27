package x

import "sync"

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
