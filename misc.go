package x

import "strings"

// Ptr returns a pointer to v.
func Ptr[T any](v T) *T {
	return &v
}

// MultiLines formats a multiline raw string, changing:
//
//	`
//		First line
//			Second line
//			Third line
//	`
//
// to:
//
//	`First line
//		Second line
//		Third line`
//
// It is intended to be called like this:
//
//	MultiLines(`
//		First Line
//			Second line
//			Third line
//	`)
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
