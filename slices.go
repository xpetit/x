package x

// Shuffle shuffles a slice randomly.
func Shuffle[S ~[]E, E any](s S) {
	n := len(s)
	if n == 0 {
		return
	}
	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	for ; i > 0; i-- {
		j := Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

// Filter filters out in place elements of the slice
func Filter[S ~[]E, E any](s S, keep func(E) bool) S {
	var i int
	for j, item := range s {
		if keep(item) {
			if i < j {
				s[i] = item
			}
			i++
		}
	}
	return s[:i]
}

func ToSet[S ~[]E, E comparable](s S) map[E]struct{} {
	m := make(map[E]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}
