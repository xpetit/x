package x

import "fmt"

// C panics if its argument is a non-nil error.
// Examples:
//
//	C(os.Chdir("directory"))
//
//	C(json.NewDecoder(os.Stdin).Decode(&data))
func C(err error) {
	if err != nil {
		panic(err)
	}
}

// C2 panics if its second argument is a non-nil error and returns the first one.
// Examples:
//
//	i := C2(strconv.Atoi("123"))
//
//	f := C2(os.Open("file"))
func C2[T any](a T, err error) T {
	C(err)
	return a
}

// C3 panics if its third argument is a non-nil error and returns the first two.
// Examples:
//
//	img, _ := C3(image.Decode(f))
//
//	_, port := C3(net.SplitHostPort(address))
func C3[T1, T2 any](a T1, b T2, err error) (T1, T2) {
	C(err)
	return a, b
}

// Assert panics if cond is false.
func Assert(cond bool, a ...any) {
	if !cond {
		if len(a) == 0 {
			panic("assertion failed")
		}
		s := "assertion failed: " + fmt.Sprintln(a...)
		panic(s[:len(s)-1])
	}
}
