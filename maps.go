package x

import (
	"cmp"
	"slices"
	"sync"
)

// Has returns whether k is present in the map m.
func Has[M ~map[K]V, K comparable, V any](m M, k K) bool {
	_, ok := m[k]
	return ok
}

func Set[M ~map[K]V, K comparable, V any](m M, k K, v V) bool {
	if _, ok := m[k]; ok {
		return true
	}
	m[k] = v
	return false
}

func Sort[S ~[]E, E cmp.Ordered](x S) S {
	slices.Sort(x)
	return x
}

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// One returns a key and value from a map, which must not be empty.
func One[M ~map[K]V, K comparable, V any](m M) (K, V) {
	for k, v := range m {
		return k, v
	}
	panic("empty map")
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	if len(m) == 0 {
		return nil
	}
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

type Protected[K comparable, V any] struct {
	M    map[K]V
	Lock sync.RWMutex
	o    sync.Once
}

func (p *Protected[K, V]) Get(k K) (v V, ok bool) {
	p.Lock.RLock()
	v, ok = p.M[k]
	p.Lock.RUnlock()
	return
}

func (p *Protected[K, _]) Has(k K) bool {
	p.Lock.RLock()
	_, ok := p.M[k]
	p.Lock.RUnlock()
	return ok
}

func (p *Protected[K, V]) Set(k K, v V) {
	p.Lock.Lock()
	p.o.Do(func() { p.M = map[K]V{} })
	p.M[k] = v
	p.Lock.Unlock()
}

func (p *Protected[K, _]) Delete(k K) {
	p.Lock.Lock()
	delete(p.M, k)
	p.Lock.Unlock()
}
