package gofuncs

import "math"

func NaN() Float {
	return math.NaN()
}

func IsPositive[T Number](value T) bool {
	return value > 0
}

func IsNotPositive[T Number](value T) bool {
	return !IsPositive(value)
}

func PanicAnyNotPositive[T Number](values ...T) {
	if value, found := IsAnyPredicate(values, IsNotPositive[T]); found {
		Panic("Only positive numbers are allowed: %v", value)
	}
}

func Abs[T SignedNumber](val T) T {
	return T(math.Abs(Float(val)))
}

func IsNegative[T SignedNumber](val T) bool {
	return math.Signbit(Float(val))
}

func SignAsInt[T SignedNumber](val T) int {
	if IsNotInit(val) || val == 0 {
		return 0
	}
	res := 1
	if IsNegative(val) {
		res *= -1
	}
	return res
}

func SignAsNumber[T SignedNumber](val T) T {
	return T(SignAsInt(val))
}

func ApplySign[T SignedNumber](isNegative bool, val T) T {
	if isNegative {
		val *= -1
	}
	return val
}

func GetIsNegativeAndAbs[T SignedNumber](val T) (bool, T) {
	return IsNegative(val), Abs(val)
}

func Trunc[T Number](number Float, precision T) Float {
	multiplier := math.Pow(10, Float(precision))
	return math.Trunc(number*multiplier) / multiplier
}

func Round[T Number](number Float, precision T) Float {
	multiplier := math.Pow(10, Float(precision))
	return math.Round(number*multiplier) / multiplier
}

func FloatToIntRound[T Int](value Float) T {
	return T(math.Round(value))
}

func Sqr[T Number](x T) T {
	return x * x
}

func Sqrt[T Number](val T) Float {
	return math.Sqrt(Float(val))
}
