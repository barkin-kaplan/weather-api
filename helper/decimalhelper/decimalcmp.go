package decimalhelper

import (
	"github.com/shopspring/decimal"
)

func AlmostEq(a, b, epsilon decimal.Decimal) bool {
	diff := a.Sub(b).Abs()
	return diff.Cmp(epsilon) <= 0
}

func Gt(a, b, epsilon decimal.Decimal) bool {
	if AlmostEq(a, b, epsilon) {
		return false
	}
	return a.GreaterThan(b)
}

func Gte(a, b, epsilon decimal.Decimal) bool {
	return a.GreaterThan(b) || AlmostEq(a, b, epsilon)
}

func Lt(a, b, epsilon decimal.Decimal) bool {
	if AlmostEq(a, b, epsilon) {
		return false
	}
	return a.LessThan(b)
}

func Lte(a, b, epsilon decimal.Decimal) bool {
	return a.LessThan(b) || AlmostEq(a, b, epsilon)
}
