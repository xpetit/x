package x

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Min returns the smaller of x or y.
func Min[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func Max[T Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Clamp[T Number](val, min, max T) T {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

// Intn returns a uniform random value in [0, n). It panics if n <= 0 or n > math.MaxInt64.
func Intn[T constraints.Integer](n T) T {
	return T(C2(rand.Int(rand.Reader, big.NewInt(int64(n)))).Int64())
}
