package x

import (
	"math"
	"strconv"
	"strings"
)

// AppendByte appends the string form of the byte count i,
// as generated by FormatByte, to dst and returns the extended buffer.
func AppendByte[T Number](dst []byte, i T) []byte {
	if i < 0 {
		panic("negative byte count")
	}
	f := float64(i)
	var thousands int
	for ; math.Round(f) >= 1000; f /= 1000 {
		thousands++
	}
	var prec int
	if thousands > 0 {
		if d := math.Round(f * 10); d < 100 && int(d)%10 != 0 {
			prec = 1
		}
	}

	dst = strconv.AppendFloat(dst, f, 'f', prec, 64)
	dst = append(dst, ' ')
	if thousands > 0 {
		dst = append(dst, " kMGTPEZY"[thousands])
	}
	dst = append(dst, 'B')
	return dst
}

func FormatByte[T Number](i T) string {
	return string(AppendByte(make([]byte, 0, 6), i))
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