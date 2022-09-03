package gofuncs

import "math"

func NaN() float64 {
	return math.NaN()
}

func IsPositive[T Number](value T) bool {
	// Nan > 0 is false
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
	return T(math.Abs(float64(val)))
}

func IsNegative[T SignedNumber](val T) bool {
	return math.Signbit(float64(val))
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

func Trunc[T Number](number float64, precision T) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Trunc(number*multiplier) / multiplier
}

func Round[T Number](number float64, precision T) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Round(number*multiplier) / multiplier
}

func FloatToIntRound[T Int](value float64) T {
	return T(math.Round(value))
}

func Sqr[T Number](x T) T {
	return x * x
}

func Sqrt[T Number](val T) float64 {
	return math.Sqrt(float64(val))
}
