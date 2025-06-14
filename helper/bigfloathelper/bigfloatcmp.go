package bigfloathelper

import (
	"math/big"
)

// AlmostEq compares a and b with the given epsilon.
// Returns true if |a - b| <= epsilon.
func AlmostEq(a, b, epsilon *big.Float) bool {
	diff := new(big.Float).Sub(a, b)
	diff.Abs(diff)
	return diff.Cmp(epsilon) <= 0
}

func Gt(a, b, epsilon *big.Float) bool {
	if AlmostEq(a, b, epsilon) {
		return false
	}
	return a.Cmp(b) > 0
}

func Gte(a, b, epsilon *big.Float) bool {
	return a.Cmp(b) > 0 || AlmostEq(a, b, epsilon)
}

func Lt(a, b, epsilon *big.Float) bool {
	if AlmostEq(a, b, epsilon) {
		return false
	}
	return a.Cmp(b) < 0
}

func Lte(a, b, epsilon *big.Float) bool {
	return a.Cmp(b) < 0 || AlmostEq(a, b, epsilon)
}
