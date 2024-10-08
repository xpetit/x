package x

import "fmt"

// Check panics if its argument is a non-nil error.
// Examples:
//
//	Check(os.Chdir("directory"))
//
//	Check(json.NewDecoder(os.Stdin).Decode(&data))
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// Must panics if its second argument is a non-nil error and returns the first one.
// Examples:
//
//	i := Must(strconv.Atoi("123"))
//
//	f := Must(os.Open("file"))
func Must[T any](a T, err error) T {
	Check(err)
	return a
}

// Must2 panics if its third argument is a non-nil error and returns the first two.
// Examples:
//
//	img, _ := Must2(image.Decode(f))
//
//	_, port := Must2(net.SplitHostPort(address))
func Must2[T1, T2 any](a T1, b T2, err error) (T1, T2) {
	Check(err)
	return a, b
}

// Assert panics if cond is false.
func Assert(cond bool, a ...any) {
	if !cond {
		if len(a) == 0 {
			panic("assertion failed")
		}
		s := "assertion failed: " + fmt.Sprintln(a...)
		s = s[:len(s)-1] // trims final newline
		panic(s)
	}
}

// AssertOK panics if ok is false. Otherwise, it returns the value.
func AssertOK[T any](v T, ok bool) T {
	if !ok {
		panic("assertion failed")
	}
	return v
}

// AssertNotNil panics if v is nil. Otherwise, it returns the value.
func AssertNotNil[T any](v *T) *T {
	if v == nil {
		panic("assertion failed")
	}
	return v
}

// AssertNonZero panics if v is a zero value. Otherwise, it returns the value.
func AssertNonZero[T comparable](v T) T {
	var z T
	if v == z {
		panic("assertion failed")
	}
	return v
}

// AssertPositive panics if v is not positive. Otherwise, it returns the value.
func AssertPositive[T Number](v T) T {
	if v > 0 {
		return v
	}
	panic("assertion failed")
}

type genericErr bool

func (genericErr) Error() string { return "generic error" }

const GenericError genericErr = false
