package x

func Check(a ...any) {
	for _, v := range a {
		if err, ok := v.(error); ok && err != nil {
			panic(err)
		}
	}
}
