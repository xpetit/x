package x

import (
	"crypto/rand"
	"math/big"
)

type (
	Integer interface {
		~uint8 | ~uint16 | ~uint32 | ~uint64 |
			~int8 | ~int16 | ~int32 | ~int64 |
			int | ~uint | ~uintptr
	}
	Float interface {
		~float32 | ~float64
	}
	Number interface {
		Integer | Float
	}
)

// Intn returns a uniform random value in [0, n). It panics if n <= 0.
func Intn[T Integer](n T) T {
	return T(Must(rand.Int(rand.Reader, big.NewInt(int64(n)))).Int64())
}

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	Must(rand.Read(b))
	return b
}
