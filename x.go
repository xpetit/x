package x

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// Check panics if its argument is a non-nil error.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// Check2 panics if its second argument is a non-nil error.
func Check2[T any](a T, err error) T {
	Check(err)
	return a
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

// Assert panics with the given message if cond is false.
func Assert(message string, cond bool) {
	if !cond {
		panic("assertion failed: " + message)
	}
}

// Min returns the smaller of x or y.
func Min[T constraints.Integer](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func Max[T constraints.Integer](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// MultiLines formats a multiline raw string, changing:
// 	`
// 		First line
// 			Second line
// 			Third line
// 	`
// to:
// 	`First line
// 		Second line
// 		Third line`
// It is intended to be called like this:
// 	MultiLines(`
// 		First Line
// 			Second line
// 			Third line
// 	`)
func MultiLines(s string) string {
	lines := strings.Split(s, "\n")
	if len(lines) < 3 {
		panic("MultiLines: expected raw string enclosed with new lines")
	}
	lines = lines[1 : len(lines)-1]
	padding := 0
loop:
	for ; padding < len(lines[0]); padding++ {
		switch lines[0][padding] {
		case '\t', '\n', '\v', '\f', '\r', ' ':
		default:
			break loop
		}
	}
	for i, line := range lines {
		lines[i] = line[Min(padding, len(line)):]
	}
	return strings.Join(lines, "\n")
}
