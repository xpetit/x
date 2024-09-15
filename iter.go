package x

import "iter"

func Limit[V any](n int, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for i := 0; i < n; i++ {
			if v, ok := next(); !ok || !yield(v) {
				return
			}
		}
	}
}

func Limit2[K, V any](n int, seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()
		for i := 0; i < n; i++ {
			if k, v, ok := next(); !ok || !yield(k, v) {
				return
			}
		}
	}
}

func Skip[V any](n int, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for i := 0; i < n; i++ {
			if _, ok := next(); !ok {
				return
			}
		}
		for {
			if v, ok := next(); !ok || !yield(v) {
				return
			}
		}
	}
}

func Skip2[K, V any](n int, seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()
		for i := 0; i < n; i++ {
			if _, _, ok := next(); !ok {
				return
			}
		}
		for {
			if k, v, ok := next(); !ok || !yield(k, v) {
				return
			}
		}
	}
}
