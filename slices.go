package x

// Reverse reverses a slice.
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

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