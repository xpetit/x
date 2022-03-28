package x

// Check panics if one of its arguments is a non-nil error.
func Check(a ...any) {
	for _, v := range a {
		if err, ok := v.(error); ok && err != nil {
			panic(err)
		}
	}
}
